package deposit

func Calculate(amount int, currency string) (int, int, string) {
	var minPercent int = 4
	var maxPercent int = 6

	minIncome := calculate(amount, minPercent)
	maxIncome := calculate(amount, maxPercent)

	return minIncome / 12, maxIncome / 12, currency

}

func calculate(depositAmount int, percent int) int {

	if percent == 0 {
		return depositAmount
	}
	return (percent * depositAmount) / 100
}
