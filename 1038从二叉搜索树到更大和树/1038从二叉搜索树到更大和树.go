// @Author:  limuye@sgepri.sgcc.com.cn
// @Desc:    二叉搜索树
// @Version: V1.0.0
// @Create:  2020/7/23/023 21:39

package LeetCode

var sum int

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func midTravel(tn *TreeNode) {
	if tn == nil {
		return
	}

	midTravel(tn.Right)
	tn.Val += sum
	sum = tn.Val
	midTravel(tn.Left)
}

func bstToGst(root *TreeNode) *TreeNode {
	midTravel(root)
	return root
}
