package entity

// Order entity
type Order struct {
	ID                   string
	CustomerName         string
	CustomerEmail        string
	CustomerAddress      string
	PaymentAmount        int64
	PaymentMethod        string
	PaymentAccountNumber string
	CartItems            []CartItem
}
