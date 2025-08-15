package wallet

type walletRepository interface {
	GetByAddress() string
}

type service struct {
	walletRepository walletRepository
}

func New(repository walletRepository) *service {
	return &service{
		walletRepository: repository,
	}
}
