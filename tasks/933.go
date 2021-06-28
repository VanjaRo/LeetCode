package main

import "fmt"

func main() {
	obj := Constructor()
	param_1 := obj.Ping(1)
	fmt.Print(param_1)
	param_2 := obj.Ping(100)
	fmt.Print(param_2)
	param_3 := obj.Ping(3001)
	fmt.Print(param_3)
	param_4 := obj.Ping(3002)
	fmt.Print(param_4)

}

type RecentCounter struct {
	pings []int
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (this *RecentCounter) Ping(t int) int {
	this.pings = append(this.pings, t)
	lowest := 0
	for this.pings[lowest] < t-3000 {
		lowest++
	}
	this.pings = this.pings[lowest:]
	return len(this.pings)
}
