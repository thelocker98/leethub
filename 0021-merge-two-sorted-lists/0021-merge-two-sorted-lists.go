/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{
		Val:  0,
		Next: nil,
	}
	joined := head

	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			joined.Next = list2
			list2 = list2.Next
		} else {
			joined.Next = list1
			list1 = list1.Next
		}
		joined = joined.Next
	}
	if list1 == nil {
		joined.Next = list2
	} else if list2 == nil {
		joined.Next = list1
	}

	return head.Next
}