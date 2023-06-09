package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"

	"github.com/zui207/CLI-sorting-visualization/sorts"
	"github.com/zui207/CLI-sorting-visualization/state"
	"github.com/zui207/CLI-sorting-visualization/visualizer"
)

func init() {
	clear()
	for i, s := range state.Collections {
		fmt.Printf("%d: %s\n", i, s)
	}
}

func setNum() int {
	cmd := exec.Command("tput", "cols")
	cmd.Stdin = os.Stdin
	out, _ := cmd.Output()
	n, _ := strconv.Atoi(strings.Trim(string(out), "\n"))
	return n
}

func input() (int, error) {
	fmt.Printf("Input 0 to %d: ", state.N-1)
	reader := bufio.NewReader(os.Stdin)
	l, _, _ := reader.ReadLine()
	s := string(l)
	return strconv.Atoi(s)
}

func exit() {
	fmt.Println("Press Any Key to Stop: ")
	reader := bufio.NewReader(os.Stdin)
	_, _, err := reader.ReadLine()
	if err == nil {
		os.Exit(0)
	}
}

func clear() {
	fmt.Fprintf(os.Stdout, "\033[2J")
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func new(arr []int) sorts.State {
	n := len(arr)
	d := make([]int, n)
	h := state.Height(n)
	p := state.Pair{A: 0, B: 0}

	return sorts.State{State: &state.State{Arr: arr, Data: [][]int{d}, Pos: []state.Pair{p}, Size: n, Height: h}}
}

func main() {
	n := setNum()
	arr := state.GenRand(n)
	s := new(arr)
	wg := sync.WaitGroup{}
	if v, err := input(); err != nil || v > state.N-1 {
		return
	} else {
		s.Id = v
		s.Sort()
		go exit()
		clear()

		wg.Add(1)
		go visualizer.Draw(s, n, &wg)
		wg.Wait()

		fmt.Fprintf(os.Stdout, "\033[%dB", s.Height+1)
		fmt.Fprintln(os.Stdout, "\x1b[?25h")
	}
}
