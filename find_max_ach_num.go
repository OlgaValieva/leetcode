package main

import "fmt"

func theMaximumAchievableX(num int, t int) int {
    x := num + t
    return x + t
}

func main() {
    num, t := 4, 1
    result := theMaximumAchievableX(num, t)
    fmt.Printf("num = %d, t = %d, Output: %d\n", num, t, result)
}

//You are given two integers, num and t.
//An integer x is called achievable if it can become equal to num after applying the following operation no more than t times:
//Increase or decrease x by 1, and simultaneously increase or decrease num by 1.
//Return the maximum possible achievable number. It can be proven that there exists at least one achievable number.