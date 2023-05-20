package state

var Collections = []string{"Insertion", "Selection", "Bubble", "Quick"}

const N int = 4

type State struct {
	Algo   string
	Id     int
	Arr    []int
	Data   [][]int
	Pos    []Pair
	Size   int
	Height int
	Count  int
}

type Pair struct {
	A int
	B int
}

func Height(n int) int {
	return (n + 7) / 8
}
