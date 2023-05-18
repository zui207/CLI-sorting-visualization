package main

func (s *State) update(arr []int, a int, b int) {
	arr[a], arr[b] = arr[b], arr[a]
	t := make([]int, s.size)
	copy(t, arr)
	s.data = append(s.data, t)
	s.pos = append(s.pos, a)
}

func (s *State) InsertionSort(arr []int) {
	for i := 1; i < s.size; i++ {
		k := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > k {
			s.update(arr, j, j+1)
			j--
		}
	}
}
