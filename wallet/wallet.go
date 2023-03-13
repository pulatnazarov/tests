package wallet

import (
	"errors"
	"fmt"
)

var (
	ErrInsufficientFunds = errors.New("not enough funds")
)

type Wallet struct {
	id      int
	balance int
}

func NewWallet(id, balance int) (Wallet, error) {
	if balance < 0 {
		return Wallet{}, fmt.Errorf("negative starting balance: %d", balance)
	}

	return Wallet{
		id:      id,
		balance: balance,
	}, nil
}

func (w *Wallet) Deposit(amount int) error {
	if amount < 0 {
		return fmt.Errorf("negative deposit amount: %d", amount)
	}

	w.balance += amount

	return nil
}

func (w *Wallet) Balance() int {
	return w.balance
}

func (w *Wallet) Withdraw(amount int) error {
	if w.balance < amount {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}
