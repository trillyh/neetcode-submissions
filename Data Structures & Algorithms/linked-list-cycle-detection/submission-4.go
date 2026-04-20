/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	curr, fast := head, head.Next

	for fast != nil && fast.Next != nil {	
		if curr == fast {
			return true
		}
		curr = curr.Next
		fast = fast.Next.Next
	}
	return false
    
}
