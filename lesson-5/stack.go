package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type Stack struct {
	elems []string
}

func (s *Stack) Len() int {
	return len(s.elems)
}

func (s *Stack) Pop() (elem string) {
	if len(s.elems) > 0 {
		elem = s.elems[len(s.elems)-1]
		s.elems = s.elems[1:5]
	}
	return elem
}

func (s *Stack) Push(elem string) {
	s.elems = append(s.elems, elem)
}

func main() {
	stack := &Stack{}
	// push "somestring"
	// pop
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		splits := strings.Split(text, " ")
		cmd := splits[0]
		switch cmd {
		case "push":
			if len(splits) > 1 {
				stack.Push(splits[1])
				log.Println("push successful; len ", stack.Len())
			} else {
				log.Println("push arg not found")
			}
		case "pop":
			if stack.Len() > 0 {
				elem := stack.Pop()
				log.Println("pop successful; elem ", elem)
			} else {
				log.Println("stack is empty - pop impossible")
			}
		}
	}
}
