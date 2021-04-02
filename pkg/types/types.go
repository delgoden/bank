package types

// Money amount of money in minimum currency units (cents, rubles, dirhams, etc.)
type Money int64

// Category the category in which the payment was made (cars, pharmacies, food, etc.)
type Category string

// Status payment status
type Status string

// Predefined payment statuses
const (
	StatusOk         Status = "OK"
	StatusFail       Status = "FAIL"
	StatusInProgress Status = "INPROGRESS"
)

// Payment payment information
type Payment struct {
	ID       int
	Amount   Money
	Category Category
	Status   Status
}
