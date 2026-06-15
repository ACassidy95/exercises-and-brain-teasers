package deletemiddlelinkednode

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode() *ListNode {
	return &ListNode{
		Next: nil,
	}
}

func (l *ListNode) AsSlice() []int {
	var vals []int = []int{}
	for l != nil {
		vals = append(vals, l.Val)
		l = l.Next
	}
	return vals
}

func DeleteMiddleNode(head *ListNode) *ListNode {
	var len, del int
	var curr, prev *ListNode

	len = 0
	for curr = head; curr != nil; curr = curr.Next {
		len++
	}

	if len > 1 {
		del = len / 2
	} else {
		del = 1
	}

	for curr = head; del != 0; curr, del = curr.Next, del-1 {
		prev = curr
	}

	if curr != nil {
		prev.Next = curr.Next
		curr.Next = nil
		curr = nil
	} else {
		head = nil
	}

	return head
}

func DeleteMiddleNodeFaster(head *ListNode) *ListNode {
	var curr, fast, prev *ListNode

	// fast moves 2 steps for curr's 1, meaning when fast becomes nil
	// curr will naturally be at the middle node
	curr, fast = head, head
	for fast.Next != nil {
		prev = curr
		curr = curr.Next
		fast = fast.Next
		if fast.Next != nil {
			fast = fast.Next
		}
	}

	if prev != nil {
		prev.Next = curr.Next
		curr.Next = nil
		curr = nil
	} else {
		head = nil
	}

	return head
}

// Floyd's Tortoise & Hare Algorithm:
// Initialises two pointers: slow and fast
// slow begins at head, and fast begins already two steps ahead.
// For 2 elements this means fast will already have gone off the end and
// we only need to delete the 2nd element (index 1)
//

func DeleteMiddleNodeOptimal(head *ListNode) *ListNode {
	var slow, fast *ListNode

	if head.Next == nil {
		return nil
	}

	slow = head
	fast = slow.Next.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	slow.Next = slow.Next.Next
	return head
}
