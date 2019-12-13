package main

import (
	"strconv"

	"github.com/kindermoumoute/adventofcode/pkg/intcode"

	"github.com/kindermoumoute/adventofcode/pkg"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {

	seq := []int{1}
	c := intcode.New(input, 0, seq...)
	part1 := strconv.Itoa(c.Run())

	seq = []int{5}
	c = intcode.New(input, 0, seq...)
	part2 := strconv.Itoa(c.Run())

	return part1, part2
}

func main() {
	pkg.Execute(run, tests, puzzle, true)
}
