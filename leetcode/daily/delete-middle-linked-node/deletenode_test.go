package deletemiddlelinkednode_test

import (
	deletemiddlelinkednode "daily/delete-middle-linked-node"
	"reflect"
	"testing"
)

func TestDeleteMiddleNode(t *testing.T) {
	testCases := map[string]struct {
		vals     []int
		expected []int
	}{
		"[1,3,4,7,1,2,6]": {
			vals:     []int{1, 3, 4, 7, 1, 2, 6},
			expected: []int{1, 3, 4, 1, 2, 6},
		},
		"[1,2,3,4]": {
			vals:     []int{1, 2, 3, 4},
			expected: []int{1, 2, 4},
		},
		"[2,1]": {
			vals:     []int{2, 1},
			expected: []int{2},
		},
		"[1]": {
			vals:     []int{1},
			expected: []int{},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testHead := createLinkedList(tc.vals)
			gotHead := deletemiddlelinkednode.DeleteMiddleNodeFaster(testHead)
			gotSlice := gotHead.AsSlice()
			if !reflect.DeepEqual(tc.expected, gotSlice) {
				t.Fatalf("Expected: %v, Got: %v", tc.expected, gotSlice)
			}
		})
	}
}

func createLinkedList(vals []int) *deletemiddlelinkednode.ListNode {
	var head, curr *deletemiddlelinkednode.ListNode

	head = deletemiddlelinkednode.NewListNode()
	curr = head
	for i, val := range vals {
		curr.Val = val
		if i < len(vals)-1 {
			curr.Next = deletemiddlelinkednode.NewListNode()
			curr = curr.Next
		}
	}

	return head
}
