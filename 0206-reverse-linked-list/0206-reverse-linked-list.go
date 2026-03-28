/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	old := &ListNode{}
	previous := head
	current := head.Next

	previous.Next = nil

	for current.Next != nil {
		old = previous
		previous = current
		current = previous.Next

		previous.Next = old
	}

	current.Next = previous

	return current
}
