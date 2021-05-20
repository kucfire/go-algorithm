/*
	leetcode tag:141 title:环形链表
*/

package point

type ListNode struct {
	Val  int
	Next *ListNode
}

// method of double points
func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}

	return false
}
