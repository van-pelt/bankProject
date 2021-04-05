package payment

import (
	"fmt"
	"github.com/van-pelt/bank/pkg/bank/card"
	"github.com/van-pelt/bank/pkg/bank/types"
)

func PaymentTesting() *[]types.Card {
	return card.GenerateCard()
}

func ExamplePaymentSources() {
	data := *PaymentTesting()
	for _, elem := range PaymentSources(data) {
		fmt.Println(elem.Number)
	}
	// Output:
	// 123 456 789
	// 123 222 789
	// 123 999 789
}
