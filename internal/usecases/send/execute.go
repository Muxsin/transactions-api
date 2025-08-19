package send

import (
	"errors"
	"infotecs-transactions-api/internal/models"

	"gorm.io/gorm"
)

var NotEnoughBalance = errors.New("not enough money to send")
var SenderWalletNotFound = errors.New("sender wallet not found")
var ReceiverWalletNotFound = errors.New("receiver wallet not found")

func (uc *UseCase) Execute(transaction *models.Transaction) error {
	senderBalance, err := uc.walletService.GetBalance(transaction.From)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return SenderWalletNotFound
		}

		return err
	}

	if *senderBalance < transaction.Amount {
		return NotEnoughBalance
	}

	uc.dbTransactionService.Begin()
	defer func() {
		if r := recover(); r != nil {
			uc.dbTransactionService.Rollback()
		}
	}()

	receiverBalance, err := uc.walletService.GetBalance(transaction.To)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ReceiverWalletNotFound
		}

		return err
	}

	if err := uc.walletService.SetBalance(transaction.From, *senderBalance-transaction.Amount); err != nil {
		uc.dbTransactionService.Rollback()
		return err
	}

	if err := uc.walletService.SetBalance(transaction.To, *receiverBalance+transaction.Amount); err != nil {
		uc.dbTransactionService.Rollback()
		return err
	}

	if err := uc.transactionService.Send(transaction); err != nil {
		uc.dbTransactionService.Rollback()
		return err
	}

	return uc.dbTransactionService.Commit()
}
