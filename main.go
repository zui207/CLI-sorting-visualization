package main

import (
	"fmt"
	"math/rand"
	"time"
)

type State struct {
	data [][]int
}

func new(arr []int) State {
	return State{data: [][]int{arr}}
}

func max() int {
	return (30 + 7) / 8
}

func genRand(n int) []int {
	arr := rand.Perm(n)
	for i := 0; i < n; i++ {
		arr[i]++
	}
	return arr
}

func main() {
	arr := genRand(30)
	a := make([]int, 30)
	copy(a, arr)
	s := new(a)
	s.InsertionSort(arr, 30)

	fmt.Println(a)
	for _, d := range s.data {
		fmt.Println("\033[2J")
		for i, n := range d {
			p := convert(i, n)
			for j := 0; j < max(); j++ {
				draw(p, n, i, j)
			}
		}
		time.Sleep(time.Millisecond * 10)
	}
}

//for _, v := range new.data {
//	fmt.Println(v)
//}

//fmt.Println("\033[2J")
//for i := 1; i < 9; i++ {
//	fmt.Printf("\033[10;%dH", i+i)
//fmt.Print(rune('▅'))
//}
//fmt.Println("\n")
//fmt.Printf("\033[%d;%dH", 0, 1)
//fmt.Println(string(int32(9608)))
//fmt.Printf("\033[%d;%dH", 2, 1)
//fmt.Println(string(int32(9608)))
