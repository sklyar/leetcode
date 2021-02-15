package c

import "math"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(node *ListNode) bool {
	for node != nil {
		if node.Val == math.MaxInt32 {
			return true
		}

		node.Val = math.MaxInt32
		node = node.Next
	}

	return false
}
