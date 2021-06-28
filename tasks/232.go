package main

import "fmt"

func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	param_2 := obj.Pop()
	param_3 := obj.Peek()
	param_4 := obj.Empty()
	fmt.Println(param_2)
	fmt.Println(param_3)
	fmt.Println(param_4)
}

type MyQueue struct {
	stack []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.stack = append(this.stack, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	ret := this.stack[0]
	this.stack = this.stack[1:]
	return ret
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.stack[0]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	lenSt := len(this.stack)
	if lenSt == 0 {
		return true
	}
	return false
}
