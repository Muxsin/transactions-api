package get_last

type transactionService interface {
	GetLast() string
}

type UseCase struct {
	transactionService transactionService
}

func New(service transactionService) *UseCase {
	return &UseCase{
		transactionService: service,
	}
}
