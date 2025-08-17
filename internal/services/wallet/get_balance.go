package wallet

func (s *service) GetBalance(address string) (*int64, error) {
	return s.walletRepository.GetByAddress(address)
}
