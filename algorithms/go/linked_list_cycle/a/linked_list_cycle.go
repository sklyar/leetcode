package a

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(node *ListNode) bool {
	nodeIndex := make(map[*ListNode]struct{})
	for {
		if node == nil {
			break
		}

		if _, ok := nodeIndex[node]; ok {
			return true
		}

		nodeIndex[node] = struct{}{}
		node = node.Next
	}

	return false
}
