package main

import (
	"testing"

	hac "github.com/dwymi02/hacunitstobase"
)


func TestHaconvertAmountToBase(t *testing.T) {
	var tests = []struct {
		input1    float64
		input2    int
		expected1 float64
	}{
		{               88.0,            244,   0.0088},
		{          1000000.0,            244, 100.0},
		{                0.0001,         254, 100.0},
		{10000000000000000.0,            234,   0.0},
		{                0.000000000001, 262,   0.0},
	}

	for _, test := range tests {
		tresult1, _ := hac.Hac_convert_amount_to_base(test.input1, test.input2)

		if tresult1 != test.expected1  {
			t.Errorf("Test Failed: {%f} {%d} inputted, {%f} expected, recieved: {%f}", test.input1, test.input2, test.expected1, tresult1)
		}
	}
}

// End of code main_test.go
