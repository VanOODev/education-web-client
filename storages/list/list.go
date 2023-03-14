package list

import (
	"errors"
	"fmt"
	"strings"
)

type List struct {
	indexes   int64
	len       int64
	firstNode *Node
	LastNode  *Node
}

func NewList() *List {
	return &List{
		indexes:   0,
		len:       0,
		firstNode: nil,
		LastNode:  nil,
	}
}

func (l *List) Add(data int64) (index int64) {
	if l.firstNode == nil {
		l.indexes += 1
		n := &Node{
			index:    l.indexes,
			data:     data,
			nextNode: nil,
		}
		l.firstNode = n
		l.LastNode = n
		l.len++
		return n.index
	}
	l.indexes += 1

	newNode := &Node{
		index:    l.indexes,
		data:     data,
		nextNode: nil,
	}

	l.LastNode.nextNode = newNode
	l.LastNode = newNode
	l.len++
	return newNode.index
}

func (l *List) Get(index int64) (data int64, err error) {
	if index <= 0 {
		err = errors.New("Index out of range")
		return
	}
	for n := l.firstNode; n != nil; n = n.nextNode {
		if n.index == index {
			return n.data, nil
		}
	}
	return
}

func (l *List) Delete(index int64) {

	if index == l.firstNode.index {
		l.firstNode = l.firstNode.nextNode
		l.len--
		return
	}

	temp1 := l.firstNode.nextNode
	temp2 := l.firstNode

	for ; temp1 != nil; temp1 = temp1.nextNode {
		if index == temp1.index {
			if temp1 == l.LastNode {
				l.LastNode = temp2
			}
			temp2.nextNode = temp1.nextNode
			return
		}
		temp2 = temp2.nextNode
	}
}

func (l *List) Len() int64 {
	return l.len
}

func (l *List) String() string {
	if l.firstNode == nil {
		return "empty list"
	}

	sb := strings.Builder{}
	n := l.firstNode
	for {
		sb.WriteString(fmt.Sprintf("index: '%d' \t data: '%d'", n.index, n.data))
		if n.nextNode == nil {
			return sb.String()
		}
		sb.WriteString("\n")
		n = n.nextNode
	}
}

func (l *List) SortIncrease() {

}

func (l *List) SortDecrease() {

}
