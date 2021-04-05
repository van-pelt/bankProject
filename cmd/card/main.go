package main

import (
	"fmt"
	"github.com/van-pelt/bank/pkg/bank/card"
	"github.com/van-pelt/bank/pkg/bank/payment"
)

func main() {

	cardsList := card.GenerateCard()
	fmt.Println(card.Total(cardsList))
	payments := payment.GeneratePayments()
	fmt.Println(payment.Max(payments))
}
