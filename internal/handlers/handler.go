package handlers

import (
	get_balance "infotecs-transactions-api/internal/usecases/get-balance"
	get_last "infotecs-transactions-api/internal/usecases/get-last"
	"infotecs-transactions-api/internal/usecases/send"
)

type handler struct {
	getBalanceUseCase get_balance.UseCase
	getLastUseCase    get_last.UseCase
	sendUseCase       send.UseCase
}

func New(
	getBalanceUseCase get_balance.UseCase,
	getLastUseCase get_last.UseCase,
	sendUseCase send.UseCase,
) *handler {
	return &handler{
		getBalanceUseCase: getBalanceUseCase,
		getLastUseCase:    getLastUseCase,
		sendUseCase:       sendUseCase,
	}
}
