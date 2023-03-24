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
			l.len--
		}
		temp2 = temp2.nextNode
	}
}

func (l *List) Len() int64 {
	return l.len
}

func (l *List) String() string {
	if l.firstNode == nil {
		return "empty l"
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
	for i := int64(1); i < l.len; i++ {
		first := l.firstNode
		second := l.firstNode.nextNode
		for j := int64(1); j < l.len; j++ {
			if first.data > second.data {
				temp1 := first.data
				temp2 := first.index
				first.index = second.index
				first.data = second.data
				second.data = temp1
				second.index = temp2

			}
			first = first.nextNode
			second = second.nextNode
		}
	}
}

func (l *List) SortIncrease1() {
	if l.len <= 1 {
		return
	}

	for i := 0; i < int(l.len)-1; i++ {
		currNode := l.firstNode
		prevNode := (*Node)(nil)
		swapped := false
		for j := 0; j < int(l.len)-i-1; j++ {
			nextNode := currNode.nextNode
			if currNode.data > nextNode.data {
				// меняем ссылки на ноды местами
				if prevNode != nil {
					prevNode.nextNode = nextNode
				} else {
					l.firstNode = nextNode
				}
				currNode.nextNode = nextNode.nextNode
				nextNode.nextNode = currNode
				currNode, nextNode = nextNode, currNode
				swapped = true
			}
			prevNode = currNode
			currNode = currNode.nextNode
		}
		// Если за проход не было перестановок, то список уже отсортирован
		if !swapped {
			break
		}
	}

	// Обновляем ссылку на последнюю ноду
	lastNode := l.firstNode
	for lastNode.nextNode != nil {
		lastNode = lastNode.nextNode
	}
	l.LastNode = lastNode
}

func (l *List) SortIncrease2() {
	if l.len <= 1 {
		return
	}

	for i := 0; i < int(l.len)-1; i++ {
		currNode := l.firstNode
		var prevNode *Node = nil
		swapped := false
		var lastSwappedNode *Node = nil // Ссылка на последнюю ноду, которая была переставлена
		for j := 0; j < int(l.len)-i-1; j++ {
			nextNode := currNode.nextNode
			if currNode.data > nextNode.data {
				// меняем ссылки на ноды местами
				if prevNode != nil {
					prevNode.nextNode = nextNode
				} else {
					l.firstNode = nextNode
				}
				currNode.nextNode = nextNode.nextNode
				nextNode.nextNode = currNode
				currNode, nextNode = nextNode, currNode
				swapped = true
				lastSwappedNode = nextNode // Сохраняем ссылку на последнюю переставленную ноду
			}
			prevNode = currNode
			currNode = currNode.nextNode
		}
		// Если за проход не было перестановок, то список уже отсортирован
		if !swapped {
			break
		}
		// Устанавливаем ссылку на последнюю переставленную ноду в качестве lastNode
		l.LastNode = lastSwappedNode
	}
}

func (l *List) SortDecrease() {

}
