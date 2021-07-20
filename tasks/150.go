package main

import (
	"fmt"
	"strconv"
)

func main() {
	arr := []string{"4", "13", "5", "/", "+"}
	fmt.Print(evalRPN(arr))
	// fmt.Print(13 / 5)
}

func evalRPN(tokens []string) int {
	st := Stack{}
	for _, token := range tokens {

		switch token {
		case "+":
			a, b := st.Pop(), st.Pop()
			st.Push(a + b)

		case "-":
			a, b := st.Pop(), st.Pop()
			st.Push(b - a)
		case "*":
			a, b := st.Pop(), st.Pop()
			st.Push(a * b)

		case "/":
			a, b := st.Pop(), st.Pop()
			st.Push(b / a)
		default:
			el, _ := strconv.Atoi(token)
			st.Push(el)

		}

	}
	return st.Pop()
}

type Stack struct {
	stack []int
}

// func NewStack() Stack {
// 	return stack
// }

func (st *Stack) Pop() int {
	ret := st.stack[len(st.stack)-1]
	st.stack = st.stack[:len(st.stack)-1]
	return ret
}

func (st *Stack) Push(el int) {
	st.stack = append(st.stack, el)
}
