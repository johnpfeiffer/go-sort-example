package main

import (
	"fmt"
	"sort"
)

// go run main.go builtin-vs-custom.go sorting-maps.go

func main() {
	People := []Person{
		{Name: "Charles", Age: 25},
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 24},
	}
	fmt.Printf("Original Listing: %v \n", People)

	// https://golang.org/pkg/sort/#SliceStable
	sort.SliceStable(People, func(i, j int) bool { return People[i].Age < People[j].Age })
	fmt.Printf("Original Sorted using BuiltIn SliceStable: %v \n", People)

	// Override the Sort interface for custom sorting behavior
	People = append(People, Person{Name: "Zorro", Age: 21})
	sort.Sort(ByAge(People))
	fmt.Printf("Original Sorted using a Customized Less: %v \n", People)

	// verbose non-interface implementation
	MySort(People)

	// examples of sorting Maps, because maps are efficient O(1) lookups but not stored in sorted order
	m := make(map[string]string)
	m["z"] = "zebra"
	m["aa"] = "Aardvark"
	m["B"] = "Bear"
	m["a"] = "antelope"
	sortedKeys := SortByKey(m)
	fmt.Printf("map sorted by Keys: ")
	for _, key := range sortedKeys {
		fmt.Printf("%s %s ", key, m[key])
	}
	fmt.Println()
	sortedKeysByValue := SortByValue(m)
	fmt.Printf("map sorted by Values: ")
	for _, key := range sortedKeysByValue {
		fmt.Printf("%s %s ", key, m[key])
	}
	fmt.Println()
}
