// @Author:  limuye@sgepri.sgcc.com.cn
// @Desc:    使用哈希表提高运行效率
// @Version: V1.0.0
// @Create:  2020/7/5 9:23

package LeetCode

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for id1, num := range nums {
		if id2, ok := numMap[target-num]; ok {
			return []int{id2, id1}
		} else {
			numMap[num] = id1
		}
	}

	return nil
}
