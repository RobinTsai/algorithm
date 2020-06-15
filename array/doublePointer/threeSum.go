/* 给定一个数组，求数组中和为 0 的三个元素
 * 返回所有可能的组合的三元组，三元组不重复且每个元组中顺序随意
 * 如：输入 `[]int{-1, 0, 1, 2, -1, -4}`，`[[-1 2 -1] [0 1 -1]]`
 */
package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{-1, 0, 1, 2, -1, -4}
	a := threeSum(arr)
	fmt.Println(a)
}
func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	result := [][]int{}
	for idx := 0; idx < len(nums); {
		target := nums[idx]
		twoSumResult := twoSum(nums[idx+1:], -target)
		for _, value := range twoSumResult {
			tempResult := append(value, target)
			result = append(result, tempResult)
		}
		idx = getNextIdx(nums, idx)
	}
	return result
}

func twoSum(nums []int, target int) [][]int {
	if len(nums) < 2 {
		return nil
	}

	result := [][]int{}
	for i, j := 0, len(nums)-1; i < j; {
		sum := nums[i] + nums[j]
		if sum > target {
			j = getLastIdx(nums, j)
		} else if sum < target {
			i = getNextIdx(nums, i)
		} else {
			result = append(result, []int{nums[i], nums[j]})
			i = getNextIdx(nums, i)
		}
	}
	return result
}

func getNextIdx(nums []int, idx int) int {
	length := len(nums)
	nextIdx := idx
	for nextIdx < length && nums[idx] == nums[nextIdx] {
		nextIdx++
	}
	return nextIdx
}

func getLastIdx(nums []int, idx int) int {
	lastIdx := idx
	for lastIdx >= 0 && nums[idx] == nums[lastIdx] {
		lastIdx--
	}
	return lastIdx
}
