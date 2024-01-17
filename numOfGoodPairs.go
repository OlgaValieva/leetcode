package main

import "fmt"

func numIdenticalPairs(nums []int) int {
    countMap := make(map[int]int)

	pairs := 0

	for _, num := range nums {
		countMap[num]++
		pairs +=countMap[num] - 1
	}
	return pairs
}

func main() {
    nums := []int{1, 2, 3, 1, 1, 3}
    result := numIdenticalPairs(nums)
    fmt.Println(result)
}