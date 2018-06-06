package linktable

import (
	"sync"
	"errors"
)

type list struct {
	FirstNode *node
	TailNode *node
	NodeCount int64

	currNode *node

	lock sync.Mutex
}

func NewLinkedList() *list {
	return &list{}
}

func (l *list) First() *node {
	return l.FirstNode
}

func (l *list) Tail() *node {
	return l.TailNode
}

func (l *list) count() int64 {
	return l.NodeCount
}

func (l *list) LPush(content interface{}) *node {
	l.lock.Lock()

	newNode := &node{Content: content}
	newNode.NextNode = l.FirstNode
	if l.NodeCount == 0 {
		l.TailNode = newNode
	}
	if l.NodeCount == 1 {
		l.TailNode = l.FirstNode
	}

	if l.FirstNode != nil {
		l.FirstNode.PreviousNode = newNode
	}
	l.FirstNode = newNode
	l.NodeCount += 1

	l.lock.Unlock()

	return newNode
}

func (l *list) RPush(content interface{}) *node {
	l.lock.Lock()

	newNode := &node{PreviousNode: l.TailNode,Content: content}

	if l.NodeCount == 0 {
		l.FirstNode = newNode
	}

	if l.TailNode != nil {
		l.TailNode.NextNode = newNode
	}

	l.TailNode = newNode
	l.NodeCount += 1

	l.lock.Unlock()

	return newNode
}

func (l *list) DeleteNode(n *node) error {
	l.lock.Lock()

	if l.NodeCount == 0 {
		errors.New("empty l")
	}

	n.PreviousNode.NextNode = n.NextNode
	n.NextNode.PreviousNode = n.PreviousNode
	l.NodeCount -= 1

	l.lock.Unlock()

	return nil
}

func (l *list) insertBefore(n *node, content interface{}) *node {
	l.lock.Lock()

	newNode := &node{PreviousNode: n.PreviousNode ,NextNode: n, Content: content}
	n.PreviousNode.NextNode = newNode
	n.PreviousNode = newNode

	l.lock.Unlock()

	return newNode
}

func (l *list) insertAfter(n *node, content interface{}) *node {
	l.lock.Lock()

	newNode := &node{PreviousNode: n, NextNode: n.NextNode ,Content: content}
	n.PreviousNode.NextNode = newNode
	n.PreviousNode = newNode

	l.lock.Unlock()

	return newNode
}

func (l *list) Next() *node {
	if l.NodeCount == 0 {
		return nil
	}

	if l.currNode == nil {
		l.currNode = l.FirstNode
	} else {
		l.currNode = l.currNode.NextNode
	}

	return l.currNode
}

func (l *list) Current() *node {
	return l.currNode
}