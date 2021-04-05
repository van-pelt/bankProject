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
