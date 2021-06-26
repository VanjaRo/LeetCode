package main

import "math/rand"

func main() {

}

type valEl struct {
	val int
	ind int
}

type RandomizedSet struct {
	vals    []int
	valsSet map[int]valEl
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{
		valsSet: make(map[int]valEl),
	}

}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.valsSet[val]; !ok {
		this.valsSet[val] = valEl{val: val, ind: len(this.vals)}
		this.vals = append(this.vals, val)
		return true
	}
	return false
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if setIt, ok := this.valsSet[val]; ok {
		last := this.vals[len(this.vals)-1]
		this.vals[setIt.ind] = last
		this.vals = this.vals[:len(this.vals)-1]
		this.valsSet[last] = valEl{last, setIt.ind}
		delete(this.valsSet, val)
		return true
	}
	return false
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	return this.vals[rand.Intn(len(this.vals))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
