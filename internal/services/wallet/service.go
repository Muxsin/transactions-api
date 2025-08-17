package wallet

type walletRepository interface {
	GetByAddress(address string) (*int64, error)
	SetBalanceByAddress(address string, balance int64) error
}

type service struct {
	walletRepository walletRepository
}

func New(repository walletRepository) *service {
	return &service{
		walletRepository: repository,
	}
}
