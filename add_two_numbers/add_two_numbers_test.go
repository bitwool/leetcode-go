package add_two_numbers

import (
	"leetcode-go/util"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	l1 := util.ArrayToLinkedList([]int{2, 4, 3})
	l2 := util.ArrayToLinkedList([]int{5, 6, 4})
	util.PrintLinkedList(addTwoNumbers(l1, l2))
}
