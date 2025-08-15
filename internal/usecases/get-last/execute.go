package get_last

func (uc *UseCase) Execute() string {
	return uc.transactionService.GetLast()
}
