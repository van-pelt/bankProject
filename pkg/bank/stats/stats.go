package stats

import "github.com/van-pelt/bank/pkg/bank/types"

func Avg(payments *[]types.Payment) types.Money {
	var sum types.Money
	var counter int
	for _, v := range *payments {
		sum += v.Amount
		counter++
	}
	if counter == 0 {
		return 0
	}
	return types.Money(int(sum) / counter)
}

func TotalInCategory(payments *[]types.Payment, category types.Category) types.Money {
	var sum types.Money
	for _, v := range *payments {
		if v.Category == category {
			sum += v.Amount
		}
	}
	return sum
}
