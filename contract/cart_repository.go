package contract

type CartRepository interface {
	CreateCart() (string, error)
}
