package transaction

type transactionRepository interface {
	GetLastByCount() string
	Create() string
}

type service struct {
	transactionRepository transactionRepository
}

func New(repository transactionRepository) *service {
	return &service{
		transactionRepository: repository,
	}
}
