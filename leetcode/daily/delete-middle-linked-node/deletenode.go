package deletemiddlelinkednode

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode() *ListNode {
	return &ListNode{}
}

func (l *ListNode) AsSlice() []int {
	var vals []int
	for l != nil {
		vals = append(vals, l.Val)
		l = l.Next
	}
	return vals
}

func DeleteMiddleNode(head *ListNode) *ListNode {
	return nil
}
