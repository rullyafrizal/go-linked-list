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

func TestLength(t *testing.T) {
	var l = LinkedList{}

	var n1 *Node = &Node{Data: 5}
	var n2 *Node = &Node{Data: 6}
	var n3 *Node = &Node{Data: 7}
	var n4 *Node = &Node{Data: 8}

	l.PushBack(n1)
	l.PushBack(n2)
	l.PushBack(n3)
	l.PushBack(n4)

	assert.Equal(t, uint(4), l.Len(), "Ukuran tidak sesuai")

	l.Pop()

	assert.Equal(t, uint(3), l.Len(), "Ukuran tidak sesuai")

	l.Push(&Node{Data: 4})

	assert.Equal(t, uint(4), l.Len(), "Ukuran tidak sesuai")
}
