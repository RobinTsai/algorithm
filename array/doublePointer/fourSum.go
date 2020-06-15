/* 求给定数组中四个元素和为 target 的所有满足条件的四元组
 * 如输入 `[1, 0, -1, 0, -2, 2]`，返回
 *  [
 *    [-1,  0, 0, 1],
 *    [-2, -1, 1, 2],
 *    [-2,  0, 0, 2]
 *  ]
 */
package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 0, -1, 0, -2, 2}
	target := 0
	result := fourSum(nums, target)
	fmt.Println(result)
}

// 有了上一题求三元组后，可以很快写出求四元组的答案（leetCode 18）
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums) // 首先要将数组排序

	result := [][]int{}
	for idx := 0; idx < len(nums); {
		threeSumResult := threeSum(nums[idx+1:], target-nums[idx])
		for _, val := range threeSumResult {
			tempResult := append(val, nums[idx])
			result = append(result, tempResult)
		}
		idx = getNextIdx(nums, idx)
	}
	return result
}

// 此为 leetCode 第 15 题，求三个数值和为 0
func threeSum(nums []int, target int) [][]int {
	// sort.Ints(nums) // 首先要将数组排序

	result := [][]int{}
	for idx := 0; idx < len(nums); {
		twoSumResult := twoSum(nums[idx+1:], target-nums[idx])
		for _, value := range twoSumResult {
			tempResult := append(value, nums[idx])
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
