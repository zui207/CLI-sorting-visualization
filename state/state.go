package state

import (
	"math/rand"
	"time"
)

var Collections = []string{"Insertion", "Selection", "Bubble", "Merge", "Quick", "Heap"}

const N int = 6

type State struct {
	Algo   string
	Id     int
	Arr    []int
	Data   [][]int
	Pos    []Pair
	Size   int
	Height int
	Count  int
}

type Pair struct {
	A int
	B int
}

func Height(n int) int {
	return (n + 7) / 8
}

func GenRand(n int) []int {
	rand.Seed(time.Now().UnixNano())
	arr := rand.Perm(n)
	for i := 0; i < n; i++ {
		arr[i]++
	}
	return arr
}
