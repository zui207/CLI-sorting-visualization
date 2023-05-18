package main

func (s *State) update(arr []int, a int, b int, n int) {
	arr[a], arr[b] = arr[b], arr[a]
	t := make([]int, n)
	copy(t, arr)
	s.data = append(s.data, t)
}

func (s *State) InsertionSort(arr []int, n int) {
	for i := 1; i < n; i++ {
		k := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > k {
			s.update(arr, j, j+1, n)
			j--
		}
	}
}
