package main

import "fmt"

type QNode struct {
	key   int
	value int
	prev  *QNode
	next  *QNode
}

func newQNode(key int, value int) *QNode {
	return &QNode{
		key:   key,
		value: value,
		prev:  nil,
		next:  nil,
	}
}

type Queue struct {
	front *QNode
	rear  *QNode
}

func (q *Queue) isEmpty() bool {
	return q.rear == nil
}

func (q *Queue) addFrontPage(key, value int) *QNode {
	page := newQNode(key, value)
	if q.front == nil && q.rear == nil {
		q.front, q.rear = page, page
	} else {
		page.next = q.front
		q.front.prev = page
		q.front = page
	}
	return page
}

func (q *Queue) moveToFront(page *QNode) {
	if page == q.front {
		return
	} else if page == q.rear {
		q.rear = q.rear.prev
		q.rear.next = nil
	} else {
		page.prev.next = page.next
		page.next.prev = page.prev
	}

	page.next = q.front
	q.front.prev = page
	q.front = page
}

func (q *Queue) removeRear() {
	if q.isEmpty() {
		return
	} else if q.front == q.rear {
		q.front, q.rear = nil, nil
	} else {
		q.rear = q.rear.prev
		q.rear.next = nil
	}
}

func (q *Queue) getRear() *QNode {
	return q.rear
}

type LRUCache struct {
	capacity int
	size     int
	pageList Queue
	pageMap  map[int]*QNode
}

func newLru(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		pageMap:  make(map[int]*QNode),
	}
}

func (lru *LRUCache) get(key int) int {
	if _, found := lru.pageMap[key]; !found {
		return -1
	}
	val := lru.pageMap[key].value
	lru.pageList.moveToFront(lru.pageMap[key])
	return val
}

func (lru *LRUCache) put(key int, value int) {
	if _, found := lru.pageMap[key]; found {
		lru.pageMap[key].value = value
		lru.pageList.moveToFront(lru.pageMap[key])
		return
	}

	if lru.size == lru.capacity {
		key := lru.pageList.getRear().key
		lru.pageList.removeRear()
		lru.size--
		delete(lru.pageMap, key)
	}
	page := lru.pageList.addFrontPage(key, value)
	lru.size++
	lru.pageMap[key] = page
}

func main() {
	cache := newLru(5)

	cache.put(3, 45)
	fmt.Println(cache.get(3))
	fmt.Println(cache.get(1))

	cache.put(1, 1)
	cache.put(1, 5)
	fmt.Println(cache.get(1))
	fmt.Println(cache.get(2))
	cache.put(8, 8)
	fmt.Println(cache.get(1))
	fmt.Println(cache.get(8))
}
