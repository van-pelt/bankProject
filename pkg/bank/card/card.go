package card

import (
	"errors"
	"github.com/van-pelt/bank/pkg/bank/types"
)

const limits = 20000
const limitsDeposit = 50000
const limitBonus = 5000

func IssueCard(currency types.Currency, color, name string) *types.Card {
	return &types.Card{
		ID:       1000,
		PAN:      "5058 1234 5678 0001",
		Balance:  0,
		Currency: currency,
		Color:    color,
		Name:     name,
		Active:   true,
	}
}

func Withdraw(card *types.Card, amount types.Money) error {
	if !card.Active {
		return errors.New("CARD IS BLOCKED")
	}
	if amount > card.Balance || amount < 0 {
		return errors.New("INVALID AMOUNT")
	}
	if amount > limits {
		return errors.New("LIMITS")
	}
	card.Balance -= amount
	return nil
}

func Deposit(card *types.Card, amount types.Money) error {
	if !card.Active {
		return errors.New("CARD IS BLOCKED")
	}
	if amount > limitsDeposit {
		return errors.New("LIMITS")
	}
	card.Balance += amount
	return nil
}

func AddBonus(card *types.Card, percent, daysInMonth, daysInYear int) error {
	if !card.Active {
		return errors.New("CARD IS BLOCKED")
	}
	if card.Balance <= 0 {
		return errors.New("CARD IS EMPTY OR NEGATIVE BALANCE")
	}
	bonus := (int(card.MinBalance) * percent * daysInMonth) / daysInYear
	if bonus > limitBonus {
		return errors.New("BONUS IS BIG")
	}
	card.Balance += types.Money(bonus)
	return nil
}

func GenerateCard() *[]types.Card {
	return &[]types.Card{
		{
			ID:         1,
			PAN:        "123 456 789",
			Balance:    1000,
			MinBalance: 300,
			Currency:   "USD",
			Color:      "Black",
			Name:       "Vasilii Pupkin",
			Active:     true,
		},
		{
			ID:         2,
			PAN:        "123 000 789",
			Balance:    0,
			MinBalance: 0,
			Currency:   "USD",
			Color:      "Black",
			Name:       "Vasilii Pupkin",
			Active:     true,
		},
		{
			ID:         3,
			PAN:        "123 222 789",
			Balance:    500,
			MinBalance: 100,
			Currency:   "USD",
			Color:      "Black",
			Name:       "Vasilii Pupkin",
			Active:     true,
		},
		{
			ID:         4,
			PAN:        "123 555 789",
			Balance:    -10,
			MinBalance: 100,
			Currency:   "USD",
			Color:      "Black",
			Name:       "Vasilii Pupkin",
			Active:     true,
		},
		{
			ID:         5,
			PAN:        "123 777 789",
			Balance:    2500,
			MinBalance: 100,
			Currency:   "USD",
			Color:      "Black",
			Name:       "Vasilii Pupkin",
			Active:     false,
		},
		{
			ID:         6,
			PAN:        "123 999 789",
			Balance:    5000,
			MinBalance: 250,
			Currency:   "USD",
			Color:      "Black",
			Name:       "Vasilii Pupkin",
			Active:     true,
		},
	}
}

func Total(cards *[]types.Card) types.Money {
	var total types.Money
	for _, val := range *cards {
		if !val.Active || val.Balance <= 0 {
			continue
		}
		total += val.Balance
	}
	return total
}
