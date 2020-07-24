// @Author:  limuye@sgepri.sgcc.com.cn
// @Desc:    二叉搜索树
// @Version: V1.0.0
// @Create:  2020/7/23/023 21:39

package LeetCode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func midTraversal(tn *TreeNode, sum *int) {
	if tn == nil {
		return
	}

	midTraversal(tn.Right, sum)
	tn.Val += *sum
	*sum = tn.Val
	midTraversal(tn.Left, sum)
}

func bstToGst(root *TreeNode) *TreeNode {
	sum := 0
	midTraversal(root, &sum)
	return root
}
