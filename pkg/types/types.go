package types

// Money amount of money in minimum currency units (cents, rubles, dirhams, etc.)
type Money int64

// Category the category in which the payment was made (cars, pharmacies, food, etc.)
type Category string

// Payment payment information
type Payment struct {
	ID       int
	Amount   Money
	Category Category
}
