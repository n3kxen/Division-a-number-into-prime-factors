package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func pirmreiz(x float64) float64 {
	var factor float64 = 2
	var result float64
	for float64(factor) <= x {
		var module float64 = math.Mod(x, factor)

		if module == 0 {
			var result = x / float64(factor)
			x = result
			if result == 0 {
				break
			}
			fmt.Printf("%d ", int(factor))
		}
		if module > 0 {
			factor += 1
		}

	}
	fmt.Println("")
	return result
}

func main() {
	var n float64
	var s string

	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Printf("Input number: ")
		_, err := fmt.Scan(&n)
		if err != nil || n < 2 {
			fmt.Println("input-output error")
		} else {
			fmt.Printf("%d: ", int(n))
			pirmreiz(n)

		}
	} else {

		for i := 1; i < len(os.Args); i++ {
			s += os.Args[i]
			num, err := strconv.Atoi(s)
			if err == nil && num > 2 {
				n = float64(num)
				fmt.Printf("%d: ", int(n))
				pirmreiz(n)
			}

			s = ""

		}
	}

}
