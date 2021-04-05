package test

import (
	"fmt"
	"github.com/van-pelt/bank/pkg/bank/transfer"
	"testing"
)

type MockData struct {
	testAmount  float64
	expectedVal float64
}

var mockDataExample = []MockData{
	{
		testAmount:  0,
		expectedVal: 0,
	},
	{
		testAmount:  500,
		expectedVal: 502.5,
	},
	{
		testAmount:  1000,
		expectedVal: 1005,
	},
}

var mockDataExampleNegative = []MockData{
	{
		testAmount:  0,
		expectedVal: 250,
	},
}

func TestTotal(t *testing.T) {

	//------------positive
	for i, pair := range mockDataExample {
		fmt.Println("Test", i, "...")
		total := transfer.Total(pair.testAmount)
		if total != pair.expectedVal {
			t.Error("For", pair.testAmount, "expected", pair.expectedVal, "got", total)
		}
	}
	//---------negative
	/*for i, pair := range mockDataExampleNegative {
		fmt.Println("Test", i, "...")
		total := transfer.Total(pair.testAmount)
		if total != pair.expectedVal {
			t.Error("For", pair.testAmount, "expected", pair.expectedVal, "got", total)
		}
	}*/
}
