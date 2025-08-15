package get_last

type walletService interface {
	GetLast() string
}

type UseCase struct {
	walletService walletService
}

func New(service walletService) *UseCase {
	return &UseCase{
		walletService: service,
	}
}
