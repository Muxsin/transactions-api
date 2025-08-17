package wallet

func (s *service) SetBalance(address string, balance int64) error {
	return s.walletRepository.SetBalanceByAddress(address, balance)
}
