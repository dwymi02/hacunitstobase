// Package hacunitstobase converts Hacash amounts from unit to base [:248]
package hacunitstobase

import (
	"errors"
)

// Function convert_amount_gt_base converts HAC amount:unit to (base) amount:248
func convert_amount_gt_base(amount float64, units int) float64 {
	iv := 0
	hac := 10.0
	hac_base := 248
	rv := 0.0

	iv = units - hac_base

	tmp := amount * hac
	iv--

	for i := 0; i < iv; i++ {
		tmp = tmp * hac
	}

	rv = tmp

	return rv
}

// Function convert_amount_lt_base converts HAC amount:unit to (base) amount:248
func convert_amount_lt_base(amount float64, units int) float64 {
	iv := 0
	hac := 10.0
	hac_base := 248
	rv := 0.0

	iv = hac_base - units

	tmp := amount / hac
	iv--

	for i := 0; i < iv; i++ {
		tmp = tmp / hac
	}

	rv = tmp

	return rv
}

// Function Hac_convert_amount_to_base determines which conversion function to utilize
func Hac_convert_amount_to_base(amount float64, units int) (float64, error) {
	rv := 0.0

	if units < 235 || units > 261 {
		return rv, errors.New("Hac_convert_amount_to_base: Units %d not supported")

	} else if units > 248 {
		rv = convert_amount_gt_base(amount, units)
	} else {
		rv = convert_amount_lt_base(amount, units)
	}

	return rv, nil
}
