package linkedlist

import "testing"

func buildList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	curr := head
	for _, v := range nums[1:] {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}
	return head
}

func listToSlice(head *ListNode) []int {
	var res []int
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

func TestAddTwoNumbers(t *testing.T) {
	l1 := buildList([]int{2, 4, 3})
	l2 := buildList([]int{5, 6, 4})

	result := AddTwoNumbers(l1, l2)
	got := listToSlice(result)
	want := []int{7, 0, 8}

	if len(got) != len(want) {
		t.Fatalf("expected length %d, got %d", len(want), len(got))
	}

	for i := range got {
		if got[i] != want[i] {
			t.Fatalf("expected %v, got %v", want, got)
		}
	}
}
