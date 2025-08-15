package get_last

func (uc *UseCase) Execute() string {
	return uc.walletService.GetLast()
}
