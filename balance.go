package main

// Balance is Load Balancing implement.
// len is Service Count
// index is Current Location.
type Balance = func(len *uint8, index *uint8) *uint8

// RoundRobin algorithm
var RoundRobin Balance = func(len *uint8, index *uint8) *uint8 {

	*index++
	return index
}
