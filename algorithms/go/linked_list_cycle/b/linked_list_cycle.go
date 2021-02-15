package b

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(node *ListNode) bool {
	slow := node
	fast := node

	for fast != nil {
		fast = fast.Next
		if fast == nil {
			break
		}

		if slow == fast {
			return true
		}

		fast = fast.Next
		slow = slow.Next
	}

	return false
}
