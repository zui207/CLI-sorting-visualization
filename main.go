package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
)

type State struct {
	arr    []int
	data   [][]int
	pos    []Pair
	size   int
	height int
	count  int
}

func new(arr []int) State {
	n := len(arr)
	a := make([]int, n)
	h := height(n)
	p := Pair{a: 0, b: 0}
	return State{arr: arr, data: [][]int{a}, size: n, height: h, pos: []Pair{p}}
}

func genRand(n int) []int {
	arr := rand.Perm(n)
	for i := 0; i < n; i++ {
		arr[i]++
	}
	return arr
}

func exit() {
	reader := bufio.NewReader(os.Stdin)
	s, _, _ := reader.ReadRune()
	if s == LF {
		os.Exit(0)
	}
}

func clear() {
	fmt.Fprintf(os.Stdout, "\033[2J")
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	n := 50
	arr := genRand(n)
	fmt.Println(arr)
	s := new(arr)

	//s.QuickSort()
	s.InsertionSort()

	go exit()
	clear()
	draw(s, n)
	fmt.Fprintf(os.Stdout, "\033[%dB", s.height)
	fmt.Println()
}
