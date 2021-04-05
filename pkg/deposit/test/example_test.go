package test

import (
	"fmt"
	"github.com/van-pelt/bank/pkg/deposit"
	"testing"
)

type MockData struct {
	testAmount     int
	expectedValMin int
	expectedValMax int
	testCurrency   string
	expectedCurr   string
}

var mockDataExample = []MockData{
	{
		testAmount:     0,
		expectedValMin: 0,
		expectedValMax: 0,
		testCurrency:   "TJS",
		expectedCurr:   "TJS",
	},
	{
		testAmount:     0,
		expectedValMin: 0,
		expectedValMax: 0,
		testCurrency:   "USD",
		expectedCurr:   "USD",
	},
	{
		testAmount:     500000,
		expectedValMin: 1666,
		expectedValMax: 2500,
		testCurrency:   "TJS",
		expectedCurr:   "TJS",
	},
	{
		testAmount:     500000,
		expectedValMin: 1666,
		expectedValMax: 2500,
		testCurrency:   "USD",
		expectedCurr:   "USD",
	},
	{
		testAmount:     1000000,
		expectedValMin: 3333,
		expectedValMax: 5000,
		testCurrency:   "TJS",
		expectedCurr:   "TJS",
	},
	{
		testAmount:     1000000,
		expectedValMin: 3333,
		expectedValMax: 5000,
		testCurrency:   "USD",
		expectedCurr:   "USD",
	},
}

var mockDataExampleNegative = []MockData{
	{
		testAmount:     500000,
		expectedValMin: 2599,
		expectedValMax: 4588,
		testCurrency:   "TJS",
		expectedCurr:   "GBP",
	},
	{
		testAmount:     0,
		expectedValMin: 10,
		expectedValMax: 100,
		testCurrency:   "USD",
		expectedCurr:   "GBP",
	},
}

func TestCalculate(t *testing.T) {

	//------------positive
	for i, pair := range mockDataExample {
		fmt.Println("Test", i, "...")
		min, max, curr := deposit.Calculate(pair.testAmount, pair.testCurrency)
		if min != pair.expectedValMin {
			t.Error("For", pair.testAmount, "expected", pair.expectedValMin, "got", min)
		}
		if max != pair.expectedValMax {
			t.Error("For", pair.testAmount, "expected", pair.expectedValMax, "got", max)
		}
		if curr != pair.expectedCurr {
			t.Error("For", pair.testCurrency, "expected", pair.expectedCurr, "got", curr)
		}
	}
	//---------negative
	/*for i, pair := range mockDataExampleNegative {
		fmt.Println("Test", i, "...")
		min, max, curr := deposit.Calculate(pair.testAmount, pair.testCurrency)
		if min != pair.expectedValMin {
			t.Error("For", pair.testAmount, "expected", pair.expectedValMin, "got", min)
		}
		if max != pair.expectedValMax {
			t.Error("For", pair.testAmount, "expected", pair.expectedValMax, "got", max)
		}
		if curr != pair.expectedCurr {
			t.Error("For", pair.testCurrency, "expected", pair.expectedCurr, "got", curr)
		}
	}*/
}
