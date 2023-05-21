package sorts

import (
	"fmt"
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

func (s *State) mergeUpdate(a int, b int, l []int, r []int) {
	i, j := 0, 0
	for k := a; k < b; k++ {
		if l[i] <= r[j] {
			s.Arr[k] = l[i]
			s.Pos = append(s.Pos, state.Pair{A: k, B: i})
			i++
		} else {
			s.Arr[k] = r[j]
			s.Pos = append(s.Pos, state.Pair{A: k, B: j})
			j++
		}
		t := make([]int, s.Size)
		copy(t, s.Arr)
		s.Data = append(s.Data, t)
		s.Count++
	}
}

// insertionsort
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

// bubblesort
func (s *State) BubbleSort() {
	for i := 0; i < s.Size-1; i++ {
		for j := 0; j < s.Size-i-1; j++ {
			if s.Arr[j] > s.Arr[j+1] {
				s.update(j, j+1)
			}
		}
	}
}

// selectionsort
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

// quicksort
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

// mergesort
func (s *State) MergeSort() {
	s._MergeSort(0, s.Size)
}

func (s *State) merge(left int, mid int, right int) {
	n1 := mid - left
	n2 := right - mid
	l := make([]int, n1+1)
	r := make([]int, n2+1)
	l[n1] = s.Size + 1
	r[n2] = s.Size + 1
	for i := 0; i < n1; i++ {
		l[i] = s.Arr[left+i]
	}
	for i := 0; i < n2; i++ {
		r[i] = s.Arr[mid+i]
	}
	s.mergeUpdate(left, right, l, r)

}

func (s *State) _MergeSort(left int, right int) {
	if right-left > 1 {
		mid := (left + right) / 2
		s._MergeSort(left, mid)
		s._MergeSort(mid, right)
		s.merge(left, mid, right)
		fmt.Println(s.Arr, left, mid, right)
	}
}
