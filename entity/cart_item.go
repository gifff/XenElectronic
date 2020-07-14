package entity

// CartItem entity
type CartItem struct {
	ID      int64
	CartID  string
	Product Product
}
