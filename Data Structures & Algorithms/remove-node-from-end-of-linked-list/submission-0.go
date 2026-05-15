/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	end := head
	curr := dummy
	for i := 0; i < n; i++ {
		end = end.Next
	}

	for end != nil {
		end = end.Next
		curr = curr.Next
	}

	temp := curr.Next
	curr.Next = curr.Next.Next
	temp.Next = nil
	return dummy.Next
}
