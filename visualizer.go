package main

import "fmt"

type Point struct {
	full     int
	fraction int
}

func convert(i int, num int) Point {
	return Point{
		full:     num / 8,
		fraction: num % 8,
	}
}

func draw(p Point, n int, i int, j int) {
	fmt.Printf("\033[%d;%dH", j+1, i+10)

	if p.fraction == 0 {
		if j >= max()-p.full {
			fmt.Println(string(9600 + int32(8)))
		}
	} else {
		if j == max()-p.full-1 {
			fmt.Println(string(9600 + int32(p.fraction)))
		} else if p.full != 0 && j >= max()-p.full {
			fmt.Println(string(9600 + int32(8)))
		}
	}

}
