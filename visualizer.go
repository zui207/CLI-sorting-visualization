package main

import (
	"fmt"
	"os"
	"time"
)

const (
	B1    rune = 9600
	B8    rune = 9608
	EMPTY rune = 32
	LF    rune = 10
)

type Buff struct {
	s []rune
}

func height(n int) int {
	return (n + 7) / 8
}

func convert(i int, num int, h int) rune {
	fraction := num % 8
	full := num / 8

	if i == h-full-1 && fraction != 0 {
		return B1 + int32(fraction)
	}
	if h-full <= i {
		return B8
	}

	return EMPTY
}

func (b *Buff) write(nums []int, pos []int, n int, i int, h int) {
	for j := 0; j < n; j++ {
		r := convert(i, nums[j], h)
		b.s = append(b.s, r)
	}
	b.s = append(b.s, LF)
}

func draw(s State, n int) {

	for _, nums := range s.data {
		b := make([]rune, 0, 110)
		buff := Buff{s: b}
		for i := 0; i < s.height; i++ {
			buff.write(nums, s.pos, n, i, s.height)
		}
		fmt.Fprintf(os.Stdout, "\r%s", string(buff.s))
		time.Sleep(time.Millisecond * 5)
		fmt.Fprintf(os.Stdout, "\033[%dA", s.height)
	}
}
