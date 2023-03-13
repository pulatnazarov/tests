package wallet

type Repository interface {
	InsertWallet(w Wallet) error
	GetWallet(id int) (Wallet, error)
	UpdateWallet(w Wallet) error
	DeleteWallet(id int) error
}
