/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
	count := 0
	countP := head
	for countP != nil {
		count++
		countP = countP.Next
	}
	fmt.Println(count)
	mid := count / 2
	if count % 2 != 0 {
		mid++
	}
	fmt.Println(mid)
	// move mid ptr to mid-1 (need to cut 1st tail off)
	midP := head
	c := 0
	for c < mid-1 {
		midP = midP.Next
		c++
	}
	tmp := midP.Next
	midP.Next = nil
	midP = tmp
	fmt.Println(midP)
	// reverse the list
	fmt.Println("reversing")
	var prev *ListNode
	for midP != nil {
		temp := midP.Next
		midP.Next = prev
		prev = midP
		midP = temp
	}	

	fmt.Println("merging")
	first, second := head, prev
	for second != nil {
		temp1, temp2 := first.Next, second.Next
		first.Next = second
		second.Next = temp1
		first, second = temp1, temp2
	} 
}
