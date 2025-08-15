package send

func (uc *UseCase) Execute() string {
	return uc.transactionService.Send()
}
