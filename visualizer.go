package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

const N int = 4
const (
	B1      rune = 9600
	B8      rune = 9608
	EMPTY   rune = 32
	LF      rune = 10
	TARGET  rune = 36
	DEFAULT rune = 39
)

var collections = []string{"Insertion", "Selection", "Bubble", "Quick"}

type Buff struct {
	s []byte
}

func height(n int) int {
	return (n + 7) / 8
}

func convert(s State, i int, t int, num int, target bool) string {
	fraction := num % 8
	full := num / 8

	var a string

	if target && s.count != t {
		a += fmt.Sprintf("\x1b[%dm", TARGET)
	} else {
		a += fmt.Sprintf("\x1b[%dm", DEFAULT)
	}

	if i == s.height-full-1 && fraction != 0 {
		a += string(B1 + int32(fraction))
	} else if s.height-full <= i {
		a += string(B8)
	} else {
		a += string(EMPTY)
	}

	return a
}

func (b *Buff) write(s State, nums []int, t int, i int) {
	for j := 0; j < s.size; j++ {
		target := j == s.pos[t].a || j == s.pos[t].b
		block := convert(s, i, t, nums[j], target)
		b.s = append(b.s, block...)
	}
	b.s = append(b.s, string(LF)...)
}

func draw(s State, n int, wg *sync.WaitGroup) {
	for t, nums := range s.data {
		b := make([]byte, 0, 1000)
		buff := Buff{s: b}
		for i := 0; i < s.height; i++ {
			buff.write(s, nums, t, i)
		}
		fmt.Fprintf(os.Stdout, "\r%s", string(buff.s))
		fmt.Printf("\x1b[%dm", DEFAULT)
		fmt.Printf("%s | SWAP: %d | Press Enter to Exit: ", s.algo, t)
		time.Sleep(time.Millisecond * 10)
		fmt.Fprintf(os.Stdout, "\033[%dA", s.height+1)
	}

	wg.Done()
}
