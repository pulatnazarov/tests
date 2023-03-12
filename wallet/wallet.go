package wallet

import "errors"

var (
	ErrInsufficientFunds = errors.New("not enough money")
)

type Wallet struct {
	balance uint
}

func (w *Wallet) Deposit(amount uint) {
	w.balance += amount
}

func (w *Wallet) Balance() uint {
	return w.balance
}

func (w *Wallet) Withdraw(amount uint) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}
