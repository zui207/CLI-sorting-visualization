package visualizer

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/zui207/CLI-sorting-visualization/sorts"
)

const (
	B1      rune = 9600
	B8      rune = 9608
	EMPTY   rune = 32
	LF      rune = 10
	TARGET  rune = 36
	DEFAULT rune = 39
)

type Buff struct {
	s []byte
}

func convert(s sorts.State, i int, t int, num int, target bool) string {
	fraction := num % 8
	full := num / 8

	var a string

	if target && s.Count != t {
		a += fmt.Sprintf("\x1b[%dm", TARGET)
	} else {
		a += fmt.Sprintf("\x1b[%dm", DEFAULT)
	}

	if i == s.Height-full-1 && fraction != 0 {
		a += string(B1 + int32(fraction))
	} else if s.Height-full <= i {
		a += string(B8)
	} else {
		a += string(EMPTY)
	}

	return a
}

func (b *Buff) write(s sorts.State, nums []int, t int, i int) {
	for j := 0; j < s.Size; j++ {
		target := j == s.Pos[t].A || j == s.Pos[t].B
		block := convert(s, i, t, nums[j], target)
		b.s = append(b.s, block...)
	}
	b.s = append(b.s, string(LF)...)
}

func Draw(s sorts.State, n int, wg *sync.WaitGroup) {
	for t, nums := range s.Data {
		b := make([]byte, 0, 1000)
		buff := Buff{s: b}
		for i := 0; i < s.Height; i++ {
			buff.write(s, nums, t, i)
		}
		fmt.Fprintf(os.Stdout, "\r%s", string(buff.s))
		fmt.Printf("\x1b[%dm", DEFAULT)
		fmt.Printf("%s | SWAP: %d | Press Enter to Exit: ", s.Algo, t)
		time.Sleep(time.Millisecond * 10)
		fmt.Fprintf(os.Stdout, "\033[%dA", s.Height+1)
	}

	wg.Done()
}
