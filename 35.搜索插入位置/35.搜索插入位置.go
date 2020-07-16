// @Author:  limuye@sgepri.sgcc.com.cn
// @Desc:    数组 二分查找
// @Version: V1.0.0
// @Create:  2020/7/5 22:07

package LeetCode

func searchInsert(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	var m int

	for l <= r {
		m = (l + r) / 2

		if nums[m] == target {
			return m
		}

		if nums[m] < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return l
}
