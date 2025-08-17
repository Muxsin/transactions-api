package get_balance

type walletService interface {
	GetBalance(address string) (*int64, error)
}

type UseCase struct {
	walletService walletService
}

func New(service walletService) *UseCase {
	return &UseCase{
		walletService: service,
	}
}
