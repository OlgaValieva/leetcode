package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		shipSizes := make(map[int]int)

		for j := 0; j < 10; j++ {
			var size int
			fmt.Scan(&size)
			shipSizes[size]++
		}

		if isValid(shipSizes) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func isValid(shipSizes map[int]int) bool {
	requiredSizes := map[int]int{1: 4, 2: 3, 3: 2, 4: 1}

	for size, count := range requiredSizes {
		if shipSizes[size] != count {
			return false
		}
	}

	return true
}