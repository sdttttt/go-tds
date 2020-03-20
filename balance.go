package main

// Balance is Load Balancing implement.
// len is Service Count
// index is Current Location.
// Index ensures that
// the position of index and index+1 is always valid in the group.
type Balance = func(len *uint8, index *uint8) uint8

// RoundRobin algorithm
var RoundRobin = func() Balance {

	return func(len *uint8, index *uint8) uint8 {
		result := *index
		*index++
		return result
	}
}
