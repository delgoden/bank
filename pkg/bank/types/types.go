package types

// Money represents the amount of money in the minimum units of currency (cents, kopecks, dirams, etc.)
type Money int64

// Category represents the category in which the payment was made (auto, pharmacy, restaurants, etc.)
type Category string

// Payment provides information about the payment
type Payment struct {
	ID       int
	Amount   Money
	Category Category
}
