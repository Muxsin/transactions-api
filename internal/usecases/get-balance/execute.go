package get_balance

func (uc *UseCase) Execute() string {
	return uc.walletService.GetBalance()
}
