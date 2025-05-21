package main

import (
	"fmt"
	"math/rand"
	"time"
)

type RandomizedSet struct {
	// values holds the elements of the set.
	values []int
	// indices maps a value to its index in values.
	indices map[int]int
}

// Constructor initializes a new RandomizedSet.
func Constructor() RandomizedSet {
	// Seed the random number generator.
	rand.Seed(time.Now().UnixNano())
	return RandomizedSet{
		values:  []int{},
		indices: make(map[int]int),
	}
}

// Insert adds val to the set if not present.
// Returns true if the element was not present, false otherwise.
func (r *RandomizedSet) Insert(val int) bool {
	if _, exists := r.indices[val]; exists {
		return false
	}
	// Append the value and record its index.
	r.values = append(r.values, val)
	r.indices[val] = len(r.values) - 1
	return true
}

// Remove deletes val from the set if present.
// Returns true if the element was present, false otherwise.
func (r *RandomizedSet) Remove(val int) bool {
	index, exists := r.indices[val]
	if !exists {
		return false
	}
	// Move the last element into the place of the element to remove.
	lastElement := r.values[len(r.values)-1]
	r.values[index] = lastElement
	r.indices[lastElement] = index

	// Remove the last element.
	r.values = r.values[:len(r.values)-1]
	delete(r.indices, val)
	return true
}

// GetRandom returns a random element from the set.
func (r *RandomizedSet) GetRandom() int {
	randomIndex := rand.Intn(len(r.values))
	return r.values[randomIndex]
}

func main() {
	rs := Constructor()
	fmt.Println(rs.Insert(1))   // true
	fmt.Println(rs.Insert(2))   // true
	fmt.Println(rs.Insert(2))   // false
	fmt.Println(rs.Remove(1))   // true
	fmt.Println(rs.Remove(3))   // false
	fmt.Println(rs.Insert(3))   // true
	fmt.Println(rs.GetRandom()) // 2 or 3
	fmt.Println(rs.GetRandom()) // 2 or 3
	fmt.Println(rs.Remove(2))   // true
	fmt.Println(rs.GetRandom())
}
