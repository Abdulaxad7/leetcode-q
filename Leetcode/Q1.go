package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := ListNode{Val: 1234}
	l2 := ListNode{Val: 4334}

	add(&l1, &l2)
}
func add(l1, l2 *ListNode) {
	var l []int

	l = append(l, l1.Val)
}
