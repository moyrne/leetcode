package no_1315_sum_of_nodes_with_even_valued_grandparent

// https://leetcode-cn.com/problems/sum-of-nodes-with-even-valued-grandparent/
/*
1315. 祖父节点值为偶数的节点和
给你一棵二叉树，请你返回满足以下条件的所有节点的值之和：
该节点的祖父节点的值为偶数。（一个节点的祖父节点是指该节点的父节点的父节点。）
如果不存在祖父节点值为偶数的节点，那么返回 0 。

示例：
输入：root = [6,7,8,2,7,1,3,9,null,1,4,null,null,null,5]
输出：18
解释：图中红色节点的祖父节点的值为偶数，蓝色节点为这些红色节点的祖父节点。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumEvenGrandparent(root *TreeNode) int {
	return dfs(root)
}

func dfs(node *TreeNode) int {
	if node == nil {
		return 0
	}

	var ans int
	// 1. value & 1 == 0， 累加子节点
	if node.Val&1 == 0 {
		ans += getChildVal(node)
	}

	// 2. 进入左子树
	ans += dfs(node.Left)

	// 3. 进入右子树
	ans += dfs(node.Right)

	return ans
}

func getChildVal(node *TreeNode) int {
	if node == nil {
		return 0
	}
	var ans int
	if node.Left != nil {
		if node.Left.Left != nil {
			ans += node.Left.Left.Val
		}
		if node.Left.Right != nil {
			ans += node.Left.Right.Val
		}
	}
	if node.Right != nil {
		if node.Right.Left != nil {
			ans += node.Right.Left.Val
		}
		if node.Right.Right != nil {
			ans += node.Right.Right.Val
		}
	}
	return ans
}
