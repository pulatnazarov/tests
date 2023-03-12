package wallet_test

import (
	"errors"
	"test/wallet"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		w := wallet.Wallet{}
		w.Deposit(10)
		assertBalance(t, w, 10)
	})

	t.Run("Withdraw", func(t *testing.T) {
		w := wallet.Wallet{}
		w.Deposit(20)
		w.Withdraw(10)

		assertBalance(t, w, 10)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		w := wallet.Wallet{}
		err := w.Withdraw(20)
		assertError(t, err, wallet.ErrInsufficientFunds)
	})
}

func assertBalance(t *testing.T, w wallet.Wallet, want uint) {
	t.Helper()

	if w.Balance() != want {
		t.Errorf("got %d, want %d", w.Balance(), want)
	}
}

func assertError(t *testing.T, err, wantErr error) {
	t.Helper()

	if !errors.Is(err, wantErr) {
		t.Errorf("got %q, wantErr %q", err.Error(), wantErr)
	}
}
