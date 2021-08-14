/***
+-----------------------------------------------------------------------------------------------------------------------------------------------+
| RE              : convert_hacunitstobase.go
|                   Example code that converts HAC value to nnn:248
| DATE            : Aug-13-2021
| AUTHOR          : MKD
| LOCATION        : SourceCode/GoLang/os/linux/Go/src/github.com/dwymi02/hacunitstobase/convert_hacunitstobase/
| BUILD STEPS     : go mod init convert_hacunitstobase
|                   go mod tidy
| RUN   STEPS     : ./convert_hacunitstobase -amount 51092992989.0 -unit 240
| OUTPUT RESULTS  : amount: 5.1092992989e+10
|                  unit: 240
|
|                  HAC 51092992989.000000:240 converted to 510.929930:248
| NOTES           : N/A
| KNOWN PROBLEMS  :	N/A
|
| CHANGE LOG      : Version       Date                                              Description of code change
|                   -------    -----------    --------------------------------------------------------------------------------------------------
|                   0.1.0      Aug-13-2021    Initial inception
+-----------------------------------------------------------------------------------------------------------------------------------------------+
*/


package main

import (
    "flag"
    "fmt"

	hac "github.com/dwymi02/hacunitstobase"
)


var (
    amount *float64
    unit   *int
)

func init() {
    amount = flag.Float64("amount", 89.0, "HAC Amount")
    unit   = flag.Int("unit", 240, "HAC Unit")
}


// Function main()
func main() {
    flag.Parse()

    tresult1, _ := hac.Hac_convert_amount_to_base(*amount, *unit)

    fmt.Println("amount:", *amount)
    fmt.Println("unit:", *unit)
    fmt.Printf("\nHAC %f:%d converted to %f:248\n", *amount, *unit, tresult1)
}

// End of code convert_hacunitstobase.go
