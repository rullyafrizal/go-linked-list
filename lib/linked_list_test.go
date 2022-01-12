package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetList(t *testing.T) {
	var l1 LinkedList = LinkedList{}

	var n1 *Node = &Node{Data: 20}
	var n2 *Node = &Node{Data: 19}
	var n3 *Node = &Node{Data: 21}

	l1.Push(n1)
	l1.Push(n2)
	l1.Push(n3)

	_, isFound := l1.Get(2)

	assert.True(t, isFound, "Failed")
}
