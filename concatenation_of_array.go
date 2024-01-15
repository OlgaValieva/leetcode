package main

import "fmt"

func getConcatenation(nums []int) []int {
    n := len(nums)
    ans := make([]int, 2*n)

    copy(ans, nums)
    copy(ans[n:], nums)

    return ans
}

func main() {
    nums := []int{7,8,9}
    fmt.Println(getConcatenation(nums))
}

//Given an integer array nums of length n, you want to create an array ans of length 2n where ans[i] == nums[i] and ans[i + n] == nums[i] for 0 <= i < n (0-indexed).
//Specifically, ans is the concatenation of two nums arrays.
//Return the array ans.