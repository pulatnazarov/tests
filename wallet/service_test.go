package wallet

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestService_OpenWallet(t *testing.T) {
	s := NewService(&MockRepo{})

	w, err := s.OpenWallet(10)
	require.NoError(t, err, "unexpected error in OpenWallet")

	assert.NotEqual(t, 0, w.id)
	assert.Equal(t, 10, w.balance)
}

func TestService_GetWallet(t *testing.T) {
	mockRepo := NewMockRepo()
	want := Wallet{
		id:      1,
		balance: 20,
	}
	mockRepo.ReturnWhenCalled("GetWallet", want)

	s := NewService(mockRepo)

	got, err := s.GetWallet(1)
	require.NoError(t, err)

	assert.Equal(t, want, got)
}

func TestService_WithdrawFunds(t *testing.T) {
	t.Run("should pass", func(t *testing.T) {
		mockRepo := NewMockRepo()
		mockRepo.ReturnWhenCalled("GetWallet", Wallet{id: 1, balance: 10})

		s := NewService(mockRepo)
		err := s.WithdrawFunds(1, 5)
		require.NoError(t, err)

		assert.Equal(t, Wallet{id: 1, balance: 5}, mockRepo.args[0])
	})

	t.Run("should fail", func(t *testing.T) {
		mockRepo := NewMockRepo()
		mockRepo.ReturnWhenCalled("GetWallet", Wallet{id: 1, balance: 10})

		s := NewService(mockRepo)
		err := s.WithdrawFunds(1, 15)

		assert.EqualError(t, err, ErrInsufficientFunds.Error())

		assert.Equal(t, 0, mockRepo.calls["UpdateWallet"], "UpdateWallet method should have not been called")
	})
}

type MockRepo struct {
	returns map[string]interface{}
	args    []Wallet
	calls   map[string]int
}

func NewMockRepo() *MockRepo {
	return &MockRepo{
		returns: make(map[string]interface{}),
		args:    make([]Wallet, 0, 1),
		calls:   make(map[string]int),
	}
}

func (r *MockRepo) ReturnWhenCalled(method string, value interface{}) {
	r.returns[method] = value
}

func (r *MockRepo) InsertWallet(w Wallet) error {
	return nil
}

func (r *MockRepo) GetWallet(id int) (Wallet, error) {
	if v, ok := r.returns["GetWallet"]; ok {
		return v.(Wallet), nil
	}

	return Wallet{}, nil
}

func (r *MockRepo) UpdateWallet(w Wallet) error {
	r.args = append(r.args, w)
	r.calls["UpdateWallet"] += 1
	return nil
}

func (r *MockRepo) DeleteWallet(id int) error {
	return nil
}
