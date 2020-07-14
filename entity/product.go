package entity

// Product entity
type Product struct {
	ID          int64
	CategoryID  int64
	Name        string
	Description string
	Photo       string
	Price       int64
}
