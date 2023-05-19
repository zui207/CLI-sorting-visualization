package main

type Pair struct {
	a int
	b int
}

func update(s *State, a int, b int) {
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
			update(s, j, j+1)
			j--
		}
	}
}

func (s *State) QuickSort() {
	_QuickSort(s, 0, s.size-1)
}

func _QuickSort(s *State, p int, r int) {
	if p < r {
		q := s.partition(p, r)
		_QuickSort(s, p, q-1)
		_QuickSort(s, q+1, r)
	}
}

func (s *State) partition(p int, r int) int {
	k := pivot(s, p, r)
	x := s.arr[k]

	i := p - 1
	for j := p; j < r+1; j++ {
		if s.arr[j] <= x {
			i++
			update(s, i, j)
			if i == k {
				k = j
			} else if j == k {
				k = i
			}
		}
	}
	update(s, i, k)

	return i
}

func pivot(s *State, p int, r int) int {
	m := (p + r) / 2
	a, b, c := s.arr[p], s.arr[m], s.arr[r]

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
