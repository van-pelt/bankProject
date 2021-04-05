package types

type Currency string
type Category string
type Money int
type Card struct {
	ID         int
	PAN        string
	Balance    Money
	MinBalance Money
	Currency   Currency
	Color      string
	Name       string
	Active     bool
}

type Payment struct {
	ID     int
	Amount Money
	Category Category
}

type PaymentSource struct {
	Type    string
	Number  string
	Balance Money
}
