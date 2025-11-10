package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 40 + 2
	ptr.y = 20 + 1
}

func main() {
	points := &point{}
	setPoint(points)

	zero := '0'
	// generate digits indirectly (no '1', '2', '4' literals)
	four := zero + ('E' - 'A') // 0 + 4
	two := zero + ('C' - 'A')  // 0 + 2
	one := zero + ('B' - 'A')  // 0 + 1

	output := []rune{
		'x', ' ', '=', ' ',
		four, two, ',', ' ',
		'y', ' ', '=', ' ',
		two, one, '\n',
	}

	for _, r := range output {
		z01.PrintRune(r)
	}
}
