package get_balance

func (uc *UseCase) Execute(address string) (*int64, error) {
	return uc.walletService.GetBalance(address)
}
