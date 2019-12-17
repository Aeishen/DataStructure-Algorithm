/*
@author : Aeishen
@data :  19/11/07, 10:16
@description :
*/

/*
题目：
	打乱一个没有重复元素的数组。

示例:
	// 以数字集合 1, 2 和 3 初始化数组。
	int[] nums = {1,2,3};
	Solution solution = new Solution(nums);

	// 打乱数组 [1,2,3] 并返回结果。任何 [1,2,3]的排列返回的概率应该相同。
	solution.shuffle();

	// 重设数组到它的初始状态[1,2,3]。
	solution.reset();

	// 随机返回数组[1,2,3]打乱后的结果。
	solution.shuffle();
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Solution struct {
	source []int
}


func Constructor(nums []int) Solution {
	return Solution{nums}
}


/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.source
}


/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	changes := []int{}
	for _,v := range this.source{
		changes = append(changes,v)
	}
	for i := len(changes) - 1;i >= 0;i--{
		one := rand.Int() % (i + 1)
		changes[i],changes[one] = changes[one],changes[i]
	}
	return changes
}

func main() {
	rand.Seed(time.Now().Unix())
	s := Constructor([]int{1,2,3})
	fmt.Println(s.Shuffle())
	fmt.Println(s.Reset())
	fmt.Println(s.Shuffle())
	fmt.Println(s.Shuffle())
}
