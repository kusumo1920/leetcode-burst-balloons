package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	nums := []int{3, 1, 5, 8}
	maxCoins := maxCoinsSolution2(nums)
	fmt.Println("max coins of", nums, "is", maxCoins)
}

func dpSolution1(nums []int, memo *map[string]int) int {
	numsInString := convertSliceToString(nums)
	if v, ok := (*memo)[numsInString]; ok {
		return v
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

func maxCoinsSolution2(nums []int) int {
	var modifiedNums []int
	modifiedNums = append(modifiedNums, append(append([]int{1}, nums...), 1)...)
	memo := make([][]int, len(nums)+2)
	for i := range memo {
		memo[i] = make([]int, len(nums)+2)
	}

	return dpSolution2(1, len(modifiedNums)-2, &modifiedNums, &memo)
}

func dpSolution2(left, right int, nums *[]int, memo *[][]int) int {
	if v := (*memo)[left][right]; v > 0 {
		return v
	}

	if right-left < 0 {
		return 0
	}

	maxCoins := 0
	for i := left; i <= right; i++ {
		// nums[i] is the last burst one
		gain := (*nums)[left-1] * (*nums)[i] * (*nums)[right+1]
		// nums[i] is fixed, recursively call left side and right side
		remaining := dpSolution2(left, i-1, nums, memo) + dpSolution2(i+1, right, nums, memo)
		maxCoins = max(maxCoins, gain+remaining)
	}

	(*memo)[left][right] = maxCoins
	return maxCoins
}
