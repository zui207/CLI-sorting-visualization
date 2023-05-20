package sorts

import (
	"reflect"

	"github.com/zui207/CLI-sorting-visualization/state"
)

type State struct {
	*state.State
}

func (s *State) Sort() {
	name := state.Collections[s.Id] + "Sort"
	s.Algo = name
	reflect.ValueOf(s).MethodByName(name).Call([]reflect.Value{})
}

func (s *State) update(a int, b int) {
	s.Arr[a], s.Arr[b] = s.Arr[b], s.Arr[a]
	t := make([]int, s.Size)
	copy(t, s.Arr)
	s.Data = append(s.Data, t)
	s.Pos = append(s.Pos, state.Pair{A: a, B: b})
	s.Count++
}

func (s *State) InsertionSort() {
	for i := 1; i < s.Size; i++ {
		k := s.Arr[i]
		j := i - 1

		for j >= 0 && s.Arr[j] > k {
			s.update(j, j+1)
			j--
		}
	}
}

func (s *State) BubbleSort() {
	for i := 0; i < s.Size-1; i++ {
		for j := 0; j < s.Size-i-1; j++ {
			if s.Arr[j] > s.Arr[j+1] {
				s.update(j, j+1)
			}
		}
	}
}

func (s *State) SelectionSort() {
	for i := 0; i < s.Size-1; i++ {
		m := i
		for j := i + 1; j < s.Size; j++ {
			if s.Arr[m] > s.Arr[j] {
				m = j
			}
		}
		if m != i {
			s.update(m, i)
		}
	}
}

func (s *State) QuickSort() {
	s._QuickSort(0, s.Size-1)
}

func (s *State) _QuickSort(p int, r int) {
	if p < r {
		q := s.partition(p, r)
		s._QuickSort(p, q-1)
		s._QuickSort(q+1, r)
	}
}

func (s *State) partition(p int, r int) int {
	k := pivot(s.Arr, p, r)
	x := s.Arr[k]

	i := p - 1
	for j := p; j < r+1; j++ {
		if s.Arr[j] <= x {
			i++
			s.update(i, j)
			if i == k {
				k = j
			} else if j == k {
				k = i
			}
		}
	}
	s.update(i, k)

	return i
}

func pivot(arr []int, p int, r int) int {
	m := (p + r) / 2
	a, b, c := arr[p], arr[m], arr[r]

	if a >= b {
		if b >= c {
			return m
		} else if c >= a {
			return p
		} else {
			return r
		}
	} else {
		if a >= c {
			return p
		} else if c >= b {
			return m
		} else {
			return r
		}
	}

}
