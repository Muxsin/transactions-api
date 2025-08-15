package send

type transactionService interface {
	Send() string
}

type UseCase struct {
	transactionService transactionService
}

func New(service transactionService) *UseCase {
	return &UseCase{
		transactionService: service,
	}
}
