package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
)

type State struct {
	data   [][]int
	size   int
	height int
	pos    []int
}

func new(arr []int) State {
	n := len(arr)
	a := make([]int, n)
	h := height(n)
	return State{data: [][]int{a}, size: n, height: h}
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
	n := 30
	arr := genRand(n)
	s := new(arr)
	s.InsertionSort(arr)

	go exit()
	clear()
	draw(s, n)
	fmt.Fprintf(os.Stdout, "\033[%dB", s.height)
	fmt.Println()

	//fmt.Print("\x1b[36m" + "abc" + "\n")

	//m := make([]byte, 0, 100)
	//for i := 0; i < 10; i++ {
	//	s := fmt.Sprintf("\x1b[%dm", i+26) + "あ" + "\n"
	//	m = append(m, s...)
	//}
	//fmt.Println(string(m))
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
