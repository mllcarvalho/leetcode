package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	ListNode1 := new(ListNode)
	ListNode2 := new(ListNode)

	ListNode1.Val = 2
	ListNode1.Next = new(ListNode)
	ListNode1.Next.Val = 4
	ListNode1.Next.Next = new(ListNode)
	ListNode1.Next.Next.Val = 3

	ListNode2.Val = 5
	ListNode2.Next = new(ListNode)
	ListNode2.Next.Val = 6
	ListNode2.Next.Next = new(ListNode)
	ListNode2.Next.Next.Val = 4

	//listNodeRes := *addTwoNumbers(ListNode1, ListNode2)

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	tmpRes := res

	for l1 != nil || l2 != nil {
		if l1 != nil {
			tmpRes.Val = tmpRes.Val + l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			tmpRes.Val = tmpRes.Val + l2.Val
			l2 = l2.Next
		}
		if tmpRes.Val/10 != 0 {
			tmpRes.Val = tmpRes.Val - 10
			tmpRes.Next = &ListNode{Val: 1}
		} else if l1 != nil || l2 != nil {
			tmpRes.Next = &ListNode{}
		}
		tmpRes = tmpRes.Next
	}
	return res

}
