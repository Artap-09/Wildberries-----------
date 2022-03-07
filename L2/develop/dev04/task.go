package main

import "strings"

func Anagram(slice []string) map[string]*[]string {
	ds := make(map[string]*disjointSets)
	result := make(map[string]*[]string)

	for _, val := range slice {
		val = strings.ToLower(val)
		set := SetName(val)
		if _, ok := ds[set]; !ok {
			ds[set] = NewDS()
			ds[set].tree = NewTree()
			ds[set].first = val
		}
		ds[set].tree.Insert(val, true)
	}

	for _, val := range ds {
		strSlice := val.tree.Println()
		if len(strSlice) > 1 {
			result[val.first] = &strSlice
		}
	}

	return result
}

type disjointSets struct {
	tree  *Tree
	first string
}

func NewDS() *disjointSets {
	return &disjointSets{}
}
