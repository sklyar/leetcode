package remove_linked_list_elements

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	for head != nil && head.Val == val {
		head = head.Next
	}

	if head == nil {
		return nil
	}

	if head.Next == nil {
		if head.Val == val {
			return nil
		}
		return head
	}

	prev := head
	current := head.Next

	for current != nil {
		if current.Val == val {
			prev.Next = current.Next
			current = current.Next
			continue
		}

		prev = current
		current = current.Next
	}

	return head
}
