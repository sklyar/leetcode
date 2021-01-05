package merge_two_sorted_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(a *ListNode, b *ListNode) *ListNode {
	var out ListNode
	current := &out

	for {
		if a == nil && b == nil {
			break
		} else if a == nil {
			current.Next = b
			break
		} else if b == nil {
			current.Next = a
			break
		}

		if a.Val < b.Val {
			current.Next = a
			current = a
			a = a.Next
		} else {
			current.Next = b
			current = b
			b = b.Next
		}
	}

	return out.Next
}
