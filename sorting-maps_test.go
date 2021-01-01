package main

import (
	"fmt"
	"reflect"
	"testing"
)

// defining the test structure separately and clear naming helps readability
type sortedMapTest struct {
	name     string
	keys     []string
	values   []string
	expected []string
}

// some repeated test cases extracted for DRY
var tcSingleLetter = sortedMapTest{name: "single", keys: []string{"a"}, values: []string{"a"},
	expected: []string{"a"}}
var tcSortedTwoChars = sortedMapTest{name: "already-sorted", keys: []string{"a", "b"}, values: []string{"alpha", "bravo"},
	expected: []string{"a", "b"}}

func TestSortByKeyEmpty(t *testing.T) {
	t.Parallel()
	m := make(map[string]string)
	actual := SortByKey(m)
	if actual != nil {
		t.Error("\nExpected:", nil, "\nReceived: ", actual)
	}

	m[""] = ""
	actual2 := SortByKey(m)
	assertSlicesEqual(t, []string{""}, actual2)
}

func TestSortByKey(t *testing.T) {
	t.Parallel()

	testCases := []sortedMapTest{
		tcSingleLetter,
		tcSortedTwoChars,
		{name: "reverse-sorted-and-caps", keys: []string{"c", "b", "A"}, values: []string{"charlie", "bravo", "alpha"},
			expected: []string{"A", "b", "c"}},
		{name: "unsorted-and-duplicate-values", keys: []string{"m", "x", "c"}, values: []string{"charlie", "charlie", "foobar"},
			expected: []string{"c", "m", "x"}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s: %v keys mapped to %v", tc.name, tc.keys, tc.values), func(t *testing.T) {
			m := createMap(tc.keys, tc.values)
			actual := SortByKey(m)
			assertSlicesEqual(t, tc.expected, actual)
		})
	}

	m := make(map[string]string)
	result := SortByKey(m)
	if result != nil {
		t.Error("\nExpected:", nil, "\nReceived: ", result)
	}
}

func TestSortByValue(t *testing.T) {
	t.Parallel()

	testCases := []sortedMapTest{
		tcSingleLetter,
		tcSortedTwoChars,
		{name: "reverse-sorted-and-caps", keys: []string{"c", "b", "A"}, values: []string{"charlie", "bravo", "alpha"},
			expected: []string{"A", "b", "c"}},
		{name: "unsorted-and-duplicate-values-stable", keys: []string{"m", "x", "c"}, values: []string{"charlie", "charlie", "foobar"},
			expected: []string{"m", "x", "c"}},
		{name: "uppercase-values", keys: []string{"1", "2", "3"}, values: []string{"Bravo", "foo", "Alpha"},
			expected: []string{"3", "1", "2"}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s: %v keys mapped to %v", tc.name, tc.keys, tc.values), func(t *testing.T) {
			m := createMap(tc.keys, tc.values)
			actual := SortByValue(m)
			assertSlicesEqual(t, tc.expected, actual)
		})
	}

	m := make(map[string]string)
	result := SortByKey(m)
	if result != nil {
		t.Error("\nExpected:", nil, "\nReceived: ", result)
	}
}

func assertSlicesEqual(t *testing.T, expected []string, result []string) {
	t.Helper()
	if !reflect.DeepEqual(expected, result) {
		t.Error("\nExpected:", expected, "\nReceived: ", result)
	}
}

func createMap(keys, values []string) map[string]string {
	m := make(map[string]string)
	for i, key := range keys {
		m[key] = values[i]
	}
	return m
}
