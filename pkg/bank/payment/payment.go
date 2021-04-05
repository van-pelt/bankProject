package payment

import (
	"github.com/van-pelt/bank/pkg/bank/types"
	"sort"
)

func GeneratePayments() *[]types.Payment {
	return &[]types.Payment{
		{
			ID:     1,
			Amount: 10,
			Category: "food",
		},
		{
			ID:     2,
			Amount: 25,
			Category: "health",
		},
		{
			ID:     3,
			Amount: 134,
			Category: "food",
		},
		{
			ID:     4,
			Amount: 8,
			Category: "health",
		},
		{
			ID:     5,
			Amount: 17,
			Category: "internet",
		},
		{
			ID:     6,
			Amount: 55,
			Category: "food",
		},
	}

}

func Max(payments *[]types.Payment) types.Payment {
	sort.Slice(*payments, func(i, j int) bool { return (*payments)[i].Amount > (*payments)[j].Amount })
	return (*payments)[0]
}

func PaymentSources(cards []types.Card) []types.PaymentSource {
	var source []types.PaymentSource
	for _, elem := range cards {
		if !elem.Active || elem.Balance <= 0 {
			continue
		}
		source = append(source, types.PaymentSource{
			Type:    "card",
			Number:  elem.PAN,
			Balance: elem.Balance,
		})
	}
	return source
}
