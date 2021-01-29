package main

import "fmt"

type Node struct {
	Next *Node
	Prev *Node
	Data int
}

type List struct {
	Len  int
	Head *Node
	Tail *Node
}

func (l *List) Find(elem int) *Node {
	if l.Head != nil {
		for tmp := l.Head; tmp.Next != nil; tmp = tmp.Next {
			if tmp.Data == elem {
				return tmp
			}
		}
	}
	return nil
}

func (l *List) Add(prev *Node, node *Node) {
	l.Len++
	if l.Head == nil {
		l.Head = node
		return
	}
	node.Next = prev.Next
	node.Prev = prev
	node.Next.Prev = node
	prev.Next = node
}

func (l *List) Append(node *Node) {
	l.Add(l.Tail, node)
}

func (l *List) Delete(node *Node) {
	node.Prev.Next = node.Next
	node.Next = nil
}

func main() {
	list := &List{}
	node := &Node{
		Data: 5,
	}
	list.Append(node)

	fmt.Println(list.Len)

	node = list.Find(5)

	list.Delete(node)

	fmt.Println(list.Len)
}
