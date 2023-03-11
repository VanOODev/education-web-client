package list

import (
	"fmt"
	"strings"
)

type List struct {
	indexes   int64
	Len       int64
	FirstNode *Node
	//LastNode  *Node
}

func NewList() *List {
	return &List{
		indexes:   0,
		Len:       0,
		FirstNode: nil,
		//LastNode:  nil,
	}
}

func (l *List) Add(data int64) (index int64) {
	if l.FirstNode == nil {
		l.indexes += 1
		n := &Node{
			index:    l.indexes,
			Data:     data,
			NextNode: nil,
		}
		l.FirstNode = n
		return n.index
	}
	n := l.FirstNode
	for n.NextNode != nil {
		n = n.NextNode
	}
	l.indexes += 1
	newNode := &Node{
		index:    l.indexes,
		Data:     data,
		NextNode: nil,
	}
	n.NextNode = newNode
	return newNode.index
}

func (l *List) String() string {
	sb := strings.Builder{}

	for n := l.FirstNode; n.NextNode != nil; n = n.NextNode {
		sb.WriteString(fmt.Sprintf("index: '%d' \t data: '%d'\n", n.index, n.Data))
	}
	return sb.String()
}