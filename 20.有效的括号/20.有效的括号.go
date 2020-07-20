// @Author:  limuye@sgepri.sgcc.com.cn
// @Desc:    栈 字符串
// @Version: V1.0.0
// @Create:  2020/7/17 22:47

package LeetCode

var cMap = map[byte]byte{')': '(', ']': '[', '}': '{'}

func isValid(s string) bool {
	b := []byte(s)
	stack := make([]byte, 0)

	for _, c := range b {
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
		} else if len(stack) == 0 {
			return false
		} else if cMap[c] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	if len(stack) != 0 {
		return false
	}

	return true
}
