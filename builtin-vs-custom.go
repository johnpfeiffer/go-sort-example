package main

import (
	"fmt"
	"sort"
)

// Person is defined by name and age
type Person struct {
	Name string
	Age  int
}

// ByAge is a sorted list of Persons
type ByAge []Person

// Len implements the interface for Sort
func (p ByAge) Len() int {
	return len(p)
}

// Swap implements the interface for Sort
func (p ByAge) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// Less implements the interface for Sort
func (p ByAge) Less(i, j int) bool {
	// customize sorting on age equivalence to use Name too
	if p[i].Age == p[j].Age {
		if p[i].Name <= p[j].Name {
			return true
		}
		return false
	}
	// otherwise ages are not equal, simply use age
	return p[i].Age < p[j].Age
}

// MySort IS NOT USING AN INTERFACE TO SORT
func MySort(People []Person) {
	// Example of manually sorting using extra primitive data structures
	AgeToPeople := make(map[int][]Person)
	for _, p := range People {
		// handle the edge case of not yet having a list entry in the map
		_, ok := AgeToPeople[p.Age]
		if !ok {
			AgeToPeople[p.Age] = []Person{p}
		} else {
			temp := AgeToPeople[p.Age]
			AgeToPeople[p.Age] = append(temp, p) // this maintains the pre-existing order on duplicates
		}
	}

	// Sort by a property and then dereference the lookup map
	// get the unique list of ages (each map key is unique)
	var Ages []int
	for age := range AgeToPeople {
		Ages = append(Ages, age)
	}
	sort.Ints(Ages)

	var PeopleSortedByAge []Person
	for _, age := range Ages {
		temp := AgeToPeople[age]
		for _, p := range temp {
			PeopleSortedByAge = append(PeopleSortedByAge, p)
		}
	}
	fmt.Printf("Sorted by age Listing: %v \n", PeopleSortedByAge)

	// Optionally create a subset of the "lowest N members"
	// this could be more efficient if it was applied during the PeopleSortedByAge loop
	PeopleSortedByAgeSubset := PeopleSortedByAge[:3]
	fmt.Printf("Stable Subset of Sorted by age: %v \n", PeopleSortedByAgeSubset)
}
