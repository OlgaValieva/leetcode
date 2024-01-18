// Требуется реализовать функцию uniqRandn, которая генерирует слайс длины n уникальных, рандомных чисел.
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    fmt.Println(uniqRandn(10))
}

func uniqRandn(n int) []int {
    if n <= 0 {
        return nil
    }
    
    rand.Seed(time.Now().UnixNano())
    
    uniqNum := make([]int, 0, n)
    uniqSet := make(map[int]struct{})
    
    for len(uniqNum) < n {
        k := rand.Intn(n * 10)
        if _,exists := uniqSet[k]; !exists {
            uniqNum = append(uniqNum, k)
            uniqSet[k] = struct{}{}
	}
}
return uniqNum
}
