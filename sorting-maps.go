package main

import "sort"

// SortByKey returns the Keys in sorted order
func SortByKey(m map[string]string) (sortedKeys []string) {
	for k := range m {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Strings(sortedKeys)
	return
}

// SortByValue returns the Keys in the sorted order of the Values (does not guarantee Stability)
func SortByValue(m map[string]string) (sortedKeys []string) {

	// create a reverse lookup mapping that can handle collisions
	// (duplicates) when the same value has multiple occurrences
	ValueToKey := make(map[string][]string)
	for k, v := range m {
		_, ok := ValueToKey[v]
		if !ok {
			ValueToKey[v] = []string{k}
		} else {
			temp := ValueToKey[v]
			ValueToKey[v] = append(temp, k)
		}
	}

	// sort the Values
	var uniqueValues []string
	for v := range ValueToKey {
		uniqueValues = append(uniqueValues, v)
	}
	sort.Strings(uniqueValues)

	// get the corresponding Keys
	for _, v := range uniqueValues {
		keys := ValueToKey[v]
		sortedKeys = append(sortedKeys, keys...)
	}

	return
}
