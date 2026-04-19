# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

class Solution:
    def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        # [0,1,2,3]
        l = None
        c = head
        while c:
            r = c.next
            c.next = l
            l = c
            c = r
        return l
        