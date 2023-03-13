package wallet

import (
	"log"
	"testing/wallet/increment"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return Service{
		repo: repo,
	}
}

func (s Service) OpenWallet(balance int) (Wallet, error) {
	w, err := NewWallet(genID(), balance)
	if err != nil {
		return Wallet{}, err
	}

	if err = s.repo.InsertWallet(w); err != nil {
		return Wallet{}, err
	}

	return w, nil
}

func (s Service) GetWallet(id int) (Wallet, error) {
	return s.repo.GetWallet(id)
}

func (s Service) WithdrawFunds(id, amount int) error {
	w, err := s.repo.GetWallet(id)
	if err != nil {
		return err
	}

	if err = w.Withdraw(amount); err != nil {
		log.Println(err)
	}

	if err = s.repo.UpdateWallet(w); err != nil {
		return err
	}

	return nil
}

func genID() int {
	return increment.Increment()
}
