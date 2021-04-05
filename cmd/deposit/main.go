package main

import (
	"flag"
	"fmt"
	"github.com/van-pelt/bank/pkg/deposit"
)

func main() {
	cmdFlagAmount := flag.Int("amount", 100000000, " a int")
	cmdFlagCurrency := flag.String("currency", "TJS", " a string")
	flag.Parse()
	if *cmdFlagAmount < 0 {
		fmt.Println("INVALID AMOUNT")
		return
	}
	if *cmdFlagCurrency != "TJS" && *cmdFlagCurrency != "USD" {
		fmt.Println("key \"-currency\" Use only USD or TJS (example -currency=\"USD\")")
		return
	}

	amount := *cmdFlagAmount
	currency := *cmdFlagCurrency
	var currName = "сомони"
	var currMinName = "дирам"
	if currency == "USD" {
		currName = "доллар/ов"
		currMinName = "цент/ов"
	}
	min, max, curr := deposit.Calculate(amount, currency)
	fmt.Println("Ежемесячный доход", curr, ",минимальный=", min, currMinName, "(", min/100, currName, ")")
	fmt.Println("Ежемесячный доход", curr, ",максимальный=", max, currMinName, "(", max/100, currName, ")")

	return
}
