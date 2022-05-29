package main

import(
	"fmt"
    // . "github.com/AllenShi/generic-modules/alg-lib/math"
    "github.com/AllenShi/generic-modules/alg-lib/math"
)

func main() {
	ints := map[string]int64{
        "first":  34,
        "second": 12,
    }

    // Initialize a map for the float values
    floats := map[string]float64{
        "first":  35.98,
        "second": 26.99,
    }

	fmt.Printf("Generic Sums: %v and %v\n",
    math.SumIntsOrFloats[string, int64](ints),
    math.SumIntsOrFloats[string, float64](floats))
}

