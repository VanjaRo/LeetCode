package main

import "fmt"

func main() {
	lru := Constructor(2)
	lru.Put(2, 1)
	lru.Put(1, 1)
	lru.Put(2, 3)
	lru.Put(4, 1)
	fmt.Println(lru.head.next)
	fmt.Println(lru.tail.prev)
	// lru.Put(2, 2)
	// // fmt.Println(lru.dLList.head)
	// lru.Get(1)
	// lru.Put(3, 3)
	// fmt.Println(lru.dLList.tail.value)
	// fmt.Println(lru.Get(2))

}

type LRUCache struct {
	cacheMap   map[int]*DLNode
	capacity   int
	head, tail *DLNode
}

type DLNode struct {
	prev  *DLNode
	next  *DLNode
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		cacheMap: make(map[int]*DLNode),
		capacity: capacity,
		head:     &DLNode{},
		tail:     &DLNode{},
	}
	lru.head.next, lru.tail.prev = lru.tail, lru.head
	return lru
}

func (this *LRUCache) Get(key int) int {
	n, ok := this.cacheMap[key]
	if !ok {
		return -1
	}
	this.delete(n)
	this.pushAfter(this.head, n)
	return n.value
}

func (this *LRUCache) Put(key int, value int) {
	n, ok := this.cacheMap[key]
	if !ok {
		if len(this.cacheMap) == this.capacity {
			endNode := this.tail.prev
			delete(this.cacheMap, endNode.key)
			this.delete(endNode)
			endNode.key, endNode.value = key, value
			n = endNode
		} else {
			n = &DLNode{
				value: value,
				key:   key,
			}
		}
		this.cacheMap[key] = n
	} else {
		n.value = value
		this.delete(n)
	}

	this.pushAfter(this.head, n)
}

func (this *LRUCache) pushAfter(current, new *DLNode) {
	next := current.next
	new.prev, new.next = current, next
	current.next, next.prev = new, new
}

func (this *LRUCache) delete(curr *DLNode) {
	prev, next := curr.prev, curr.next
	prev.next, next.prev = next, prev
}
