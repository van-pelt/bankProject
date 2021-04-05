package stats

import (
	"fmt"
	"github.com/van-pelt/bank/pkg/bank/payment"
)

func ExampleAvg() {
	data := Avg(payment.GeneratePayments())
	fmt.Println(data)
	// Output:
	// 57
}

func ExampleTotalInCategory() {
	data := TotalInCategory(payment.GeneratePayments(), "food")
	fmt.Println(data)
	// Output:
	// 199
}
