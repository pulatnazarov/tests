package wallet_test

import (
	"errors"
	"testing"
	"testing/wallet"

	"github.com/stretchr/testify/assert"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		w := wallet.Wallet{}
		w.Deposit(10)

		assert.Equal(t, int(10), w.Balance(), "balance does not match!")
	})

	t.Run("withdraw", func(t *testing.T) {
		w := wallet.Wallet{}
		w.Deposit(20)
		w.Withdraw(10)

		assertBalance(t, w, 10)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		w := wallet.Wallet{}
		err := w.Withdraw(20)

		assertError(t, err, wallet.ErrInsufficientFunds)
	})
}

func assertBalance(t *testing.T, w wallet.Wallet, want int) {
	t.Helper()

	if w.Balance() != want {
		t.Errorf("want %d, got %d", want, w.Balance())
	}
}

func assertError(t *testing.T, err error, wantErr error) {
	t.Helper()

	if !errors.Is(err, wantErr) {
		t.Errorf("want error %q but got %q", wantErr, err)
	}
}
