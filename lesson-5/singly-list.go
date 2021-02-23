package main

import (
	"bytes"
)

type SNode struct {
	data []byte
	next *SNode
}

func NewSNode() *SNode {
	return &SNode{
		data: nil,
		next: nil,
	}
}

func (n *SNode) SetData(data []byte) {
	n.data = data
}

func (n *SNode) GetData() []byte {
	return n.data
}

func (n *SNode) SetNext(next *SNode) {
	n.next = next
}

func (n *SNode) GetNext() *SNode {
	return n.next
}

type SingleList struct {
	head *SNode
}

func NewSingleList() *SingleList {
	return &SingleList{
		head: nil,
	}
}

func (s *SingleList) InsertAtHead(data []byte) {
	node := NewSNode()
	node.SetData(data)

	if s.head != nil {
		node.SetNext(s.head)
		s.head = node
	} else {
		s.head = node
	}
}

func (s *SingleList) InsertAtTail(data []byte) {
	node := NewSNode()
	node.SetData(data)

	if s.head != nil {
		cur := s.head
		for cur.GetNext() != nil {
			cur = cur.GetNext()
		}
		cur.SetNext(node)
	} else {
		s.head = node
	}
}

func (s *SingleList) InsertAfterTarget(data []byte, target *SNode) {
	if s.head != nil {
		cur := s.head
		for cur.GetNext() != target && cur.GetNext() != nil {
			cur = cur.GetNext()
		}
		if cur == target {
			node := NewSNode()
			node.SetData(data)

			next := cur.GetNext()
			cur.SetNext(node)
			node.SetNext(next)
		}
	}
}

func (s *SingleList) Delete(target *SNode) bool {
	if s.head == nil {
		return false
	}

	cur := s.head
	var prev *SNode
	prev = nil

	for cur.GetNext() != target && cur.GetNext() != nil {
		prev = cur
		cur = cur.GetNext()
	}

	if cur == target {
		if prev != nil {
			prev.SetNext(cur.GetNext())
		} else {
			s.head = cur.GetNext()
		}
		cur = nil
		return true
	}

	return false
}

func (s *SingleList) Find(data []byte) *SNode {
	if s.head == nil {
		for cur := s.head; cur.GetNext() != nil; cur = cur.GetNext() {
			res := bytes.Compare(cur.GetData(), data)
			if res == 0 {
				return cur
			}
		}
	}

	return nil
}

func main() {
	single := NewSingleList()
	data1 := []byte{'h', 'e', 'l', 'l', 'o'}
	data2 := []byte{'g', 'o', 'l', 'a', 'n', 'g'}

	single.InsertAtHead(data1)

	node := single.Find(data1)
	single.InsertAfterTarget(data2, node)

	single.Delete(node)
}
