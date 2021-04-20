/*
 * @Author: LongZhiQiu
 * @Date: 2020-01-17 08:32:48
 * @Version: Do not edit
 * @Description: file content
 * @LastEditors  : LongZhiQiu
 * @LastEditTime : 2020-01-17 09:55:24
 * @Auditor: LongZhiQiu
 */

// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
// 你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
// 示例:
// 给定 nums = [2, 7, 11, 15], target = 9
// 因为 nums[0] + nums[1] = 2 + 7 = 9
// 所以返回 [0, 1]

package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
}

func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for index, value := range nums {
		k, ok := hashTable[value]
		if ok {
			return []int{index, k}
		}
		hashTable[target-value] = index
	}
	return []int{}
}
