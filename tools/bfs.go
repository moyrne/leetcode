package tools

type ListNode struct {
	Val         int
	Left, Right *ListNode
}

type state struct {
	node *ListNode
	deep int
}

func BFS(root *ListNode) {
	if root == nil {
		return
	}

	q := []*state{{node: root, deep: 1}}

	for len(q) != 0 {
		sz := len(q)

		for i := 0; i < sz; i++ {
			cur := q[0]
			q = q[1:]

			// do somethin

			if cur.node.Left != nil {
				q = append(q, &state{node: cur.node.Left, deep: cur.deep + 1})
			}
			if cur.node.Right != nil {
				q = append(q, &state{node: cur.node.Right, deep: cur.deep + 1})
			}
		}
	}
}
