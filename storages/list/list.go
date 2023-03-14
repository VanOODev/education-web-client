package list

import (
	"fmt"
	"log"
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

func (l *List) Get(index int64) (data int64) {
	if index < l.firstNode.index || index > l.LastNode.index {
		log.Fatal("Out of List")
	}
	temp := l.firstNode
	for i := int64(1); i < index; i++ {
		temp = temp.nextNode
	}
	return temp.data
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
