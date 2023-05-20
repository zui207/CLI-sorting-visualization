package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"sync"
	"time"
)

type State struct {
	algo   string
	id     int
	arr    []int
	data   [][]int
	pos    []Pair
	size   int
	height int
	count  int
}

func init() {
	clear()
	for i, s := range collections {
		fmt.Printf("%d: %s\n", i, s)
	}
}

func input() (int, error) {
	fmt.Printf("Input 0 to %d: ", N-1)
	reader := bufio.NewReader(os.Stdin)
	l, _, _ := reader.ReadLine()
	s := string(l)
	return strconv.Atoi(s)
}

func new(arr []int) State {
	n := len(arr)
	a := make([]int, n)
	h := height(n)
	p := Pair{a: 0, b: 0}
	return State{arr: arr, data: [][]int{a}, size: n, height: h, pos: []Pair{p}}
}

func genRand(n int) []int {
	rand.Seed(time.Now().UnixNano())
	arr := rand.Perm(n)
	for i := 0; i < n; i++ {
		arr[i]++
	}
	return arr
}

func exit() {
	fmt.Println("Press Enter to Stop: ")
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
	n := 100
	arr := genRand(n)
	s := new(arr)
	wg := sync.WaitGroup{}

	if v, err := input(); err != nil || v > N-1 {
		return
	} else {
		s.id = v
		s.Sort()
		go exit()
		clear()
		wg.Add(1)
		go draw(s, n, &wg)
		wg.Wait()

		fmt.Fprintf(os.Stdout, "\033[%dB", s.height+1)
		fmt.Println()
	}
}
