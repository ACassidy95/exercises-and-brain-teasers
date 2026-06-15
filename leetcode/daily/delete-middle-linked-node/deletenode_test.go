package deletemiddlelinkednode_test

import (
	deletemiddlelinkednode "daily/delete-middle-linked-node"
	"testing"
)

func TestDeleteMiddleNode(t *testing.T) {

}

func createLinkedList(vals []int) *deletemiddlelinkednode.ListNode {
	var head, curr *deletemiddlelinkednode.ListNode

	head = deletemiddlelinkednode.NewListNode()
	curr = head
	for _, val := range vals {
		curr.Val = val
		curr.Next = deletemiddlelinkednode.NewListNode()
		curr = curr.Next
	}

	return head
}
