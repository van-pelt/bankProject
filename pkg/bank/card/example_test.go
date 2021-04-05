package card

import (
	"fmt"
	"github.com/van-pelt/bank/pkg/bank/types"
)

func CardTesting(amount types.Money, isActive bool) *types.Card {
	MyCard := IssueCard("TJS", "BLACK", "MyBigCard")
	MyCard.Balance = amount
	MyCard.Active = isActive
	return MyCard
}

func ExampleWithdraw_positive() {
	card := CardTesting(10000, true)
	err := Withdraw(card, 5000)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(card.Balance)
	// Output:
	// 5000
}

func ExampleWithdraw_noMoney() {
	card := CardTesting(0, true)
	err := Withdraw(card, 5000)
	if err != nil {
		fmt.Println(err.Error())
		// Output:
		// INVALID AMOUNT
		return
	}
	fmt.Println(card.Balance)

}

func ExampleWithdraw_inactive() {
	card := CardTesting(10000, false)
	err := Withdraw(card, 5000)
	if err != nil {
		fmt.Println(err.Error())
		// Output:
		// CARD IS BLOCKED
		return
	}
	fmt.Println(card.Balance)

}

func ExampleWithdraw_limit() {
	card := CardTesting(100000, true)
	err := Withdraw(card, 35000)
	if err != nil {
		fmt.Println(err.Error())
		// Output:
		// LIMITS
		return
	}
	fmt.Println(card.Balance)
}

//------------------------------------------------------------------------------------------
func ExampleDeposit_normal() {
	card := CardTesting(-1000, true)
	err := Deposit(card, 2000)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(card.Balance)
	// Output:
	// 1000
}

func ExampleDeposit_blocked() {
	card := CardTesting(1000, false)
	err := Deposit(card, 2000)
	if err != nil {
		fmt.Println(err.Error())
		// Output:
		// CARD IS BLOCKED
		return
	}
	fmt.Println(card.Balance)
}

func ExampleDeposit_limit() {
	card := CardTesting(1000, true)
	err := Deposit(card, 70000)
	if err != nil {
		fmt.Println(err.Error())
		// Output:
		// LIMITS
		return
	}
	fmt.Println(card.Balance)
}

//------------------------------------------------------------------------

func ExampleAddBonus_normal() {
	card := CardTesting(10000, true)
	card.MinBalance = 10000
	err := AddBonus(card, 3, 30, 365)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(card.Balance)
	// Output:
	// 12465
}

func ExampleAddBonus_blocked() {
	card := CardTesting(10000, false)
	card.MinBalance = 10000
	err := AddBonus(card, 3, 30, 365)
	if err != nil {
		fmt.Println(err.Error())
		// Output:
		// CARD IS BLOCKED
		return
	}
	fmt.Println(card.Balance)
}

func ExampleAddBonus_negative() {
	card := CardTesting(-100, true)
	card.MinBalance = 10000
	err := AddBonus(card, 3, 30, 365)
	if err != nil {
		fmt.Println(err.Error())
		// Output:
		// CARD IS EMPTY OR NEGATIVE BALANCE
		return
	}
	fmt.Println(card.Balance)
}

func ExampleAddBonus_overbonus() {
	card := CardTesting(100, true)
	card.MinBalance = 10000
	err := AddBonus(card, 100, 30, 365)
	if err != nil {
		fmt.Println(err.Error())
		// Output:
		// BONUS IS BIG
		return
	}
	fmt.Println(card.Balance)
}
