package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	nums := []int{3, 1, 5, 8}
	maxCoins := maxCoinsSolution1(nums)
	fmt.Println("max coins of", nums, "is", maxCoins)
}

func dpSolution1(nums []int, memo *map[string]int) int {
	numsInString := convertSliceToString(nums)
	if _, ok := (*memo)[numsInString]; ok {
		return (*memo)[numsInString]
	}

	if len(nums) == 2 {
		return 0
	}

	maxCoins := 0
	for i := 1; i <= len(nums)-2; i++ {
		gain := nums[i-1] * nums[i] * nums[i+1]
		numsWithoutI := make([]int, 0)
		numsWithoutI = append(numsWithoutI, nums[:i]...)
		numsWithoutI = append(numsWithoutI, nums[i+1:]...)
		remaining := dpSolution1(numsWithoutI, memo)
		maxCoins = maxInt(maxCoins, gain+remaining)
	}

	(*memo)[numsInString] = maxCoins
	return maxCoins
}

func maxCoinsSolution1(nums []int) int {
	modifiedNums := []int{1}
	modifiedNums = append(modifiedNums, nums...)
	modifiedNums = append(modifiedNums, 1)
	emptyMemo := make(map[string]int)

	return dpSolution1(modifiedNums, &emptyMemo)
}

func convertSliceToString(nums []int) string {
	stringList := make([]string, len(nums))
	for i, v := range nums {
		stringList[i] = strconv.Itoa(v)
	}

	return strings.Join(stringList, ",")
}

func maxInt(num1, num2 int) int {
	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}
