package linkedlist

/*
You are given two non-empty linked lists representing two non-negative integers.
The digits are stored in reverse order, and each of their nodes contains a single digit.

Add the two numbers and return the sum as a linked list.
You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:
Input: l1 = [2,4,3], l2 = [5,6,4]
Explanation:
342 + 465 = 807
Output: [7,0,8]

Explanation of output:
The digits are stored reversed:
2 -> 4 -> 3  represents 342
5 -> 6 -> 4  represents 465

342 + 465 = 807

Result (reversed): 7 -> 0 -> 8

Constraints:
- The number of nodes in each linked list is in the range [1, 100].
- 0 <= Node.Val <= 9
- The lists contain non-negative integers, no leading zeros.

Solution:
Time Complexity:  O(max(n, m))
Space Complexity: O(max(n, m))
*/

// ListNode represents a node in a singly linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// AddTwoNumbers adds two numbers represented by reversed linked lists
// and returns the result as a linked list in reversed order.
// Time Complexity:  O(max(n, m))
// Space Complexity: O(max(n, m))
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: 0}
	curr := head
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		var v1, v2 int

		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		_sum := v1 + v2 + carry
		digit := _sum % 10
		carry = _sum / 10

		curr.Next = &ListNode{Val: digit}
		curr = curr.Next
	}

	return head.Next
}
