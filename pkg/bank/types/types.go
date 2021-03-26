package types

// Money gives type for money
type Money int64

// Currency gives type for currency
type Currency string

// Currency codes
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

// PAN gives type for PAN
type PAN string

// Card creates card model
type Card struct {
	ID         int
	PAN        PAN
	Balance    Money
	MinBalance Money
	Currency   Currency
	Color      string
	Name       string
	Active     bool
}

// Payment info for payments
type Payment struct {
	ID     int
	Amount Money
}

// PaymentSources ways to make a payment
type PaymentSources struct {
	Type    string
	Number  PAN
	Balance Money
}
