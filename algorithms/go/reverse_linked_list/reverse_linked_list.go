package reverse_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var (
		result  *ListNode
		current = head
	)

	for current != nil {
		next := current.Next
		current.Next = result
		result = current
		current = next
	}

	return result
}
