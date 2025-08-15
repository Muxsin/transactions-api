package get_balance

type walletService interface {
	GetBalance() string
}

type UseCase struct {
	walletService walletService
}

func New(service walletService) *UseCase {
	return &UseCase{
		walletService: service,
	}
}
