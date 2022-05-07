package homework

// package main

import (
	"sort"
)

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

func sortMapValues(input map[int]string) (result []string) {
	//Place your code here
	s := make(PairList, len(input))

	i := 0
	for k, v := range input {
		s[i] = Pair{v, k}
		i++
	}

	sort.Sort(s)

	out := make([]string, len(input))
	i = 0
	for _, k := range s {
		out[i] = k.Key
		i++
	}

	return out
}

// func main() {

// 	var m map[int]string
// 	m = make(map[int]string)
// 	m[0] = "zero"
// 	m[4] = "four"
// 	m[3] = "three"
// 	m[2] = "two"
// 	m[1] = "one"
// 	m[9] = "nine"
// 	m[8] = "eight"
// 	m[7] = "seven"
// 	m[6] = "six"
// 	m[5] = "five"

// 	for i, k := range m {
// 		fmt.Println("Key:", k, "Value:", i)
// 	}
// 	fmt.Println(len(m))
// 	fmt.Println(sortMapValues(m))
// }
