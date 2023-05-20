package main

import "reflect"

type Pair struct {
	a int
	b int
}

func (s *State) Sort() {
	name := collections[s.id] + "Sort"
	s.algo = name
	reflect.ValueOf(s).MethodByName(name).Call([]reflect.Value{})
}

func (s *State) update(a int, b int) {
	s.arr[a], s.arr[b] = s.arr[b], s.arr[a]
	t := make([]int, s.size)
	copy(t, s.arr)
	s.data = append(s.data, t)
	s.pos = append(s.pos, Pair{a: a, b: b})
	s.count++
}

func (s *State) InsertionSort() {
	for i := 1; i < s.size; i++ {
		k := s.arr[i]
		j := i - 1

		for j >= 0 && s.arr[j] > k {
			s.update(j, j+1)
			j--
		}
	}
}

func (s *State) BubbleSort() {
	for i := 0; i < s.size-1; i++ {
		for j := 0; j < s.size-i-1; j++ {
			if s.arr[j] > s.arr[j+1] {
				s.update(j, j+1)
			}
		}
	}
}

func (s *State) SelectionSort() {
	for i := 0; i < s.size-1; i++ {
		m := i
		for j := i + 1; j < s.size; j++ {
			if s.arr[m] > s.arr[j] {
				m = j
			}
		}
		if m != i {
			s.update(m, i)
		}
	}
}

func (s *State) QuickSort() {
	s._QuickSort(0, s.size-1)
}

func (s *State) _QuickSort(p int, r int) {
	if p < r {
		q := s.partition(p, r)
		s._QuickSort(p, q-1)
		s._QuickSort(q+1, r)
	}
}

func (s *State) partition(p int, r int) int {
	k := pivot(s.arr, p, r)
	x := s.arr[k]

	i := p - 1
	for j := p; j < r+1; j++ {
		if s.arr[j] <= x {
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
