package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Trees struct {
	right int
	down int
	y int
	x int
	num int
}

func (t *Trees) calcSlope(area []string) int {
	width := len(area[0])
	for {
		if t.y > len(area)-1 {
			break
		}
		line := area[t.y]
		c := line[t.x%width]
		if c == '#' {
			t.num++
		}
		t.x += t.right
		t.y += t.down
	}
	return t.num
}


func main() {
	content, err := ioutil.ReadFile("day-3-toboggan-trajectory/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	split := strings.Split(strings.TrimRight(string(content), "\n"), "\n")

	trees := Trees{
		x: 0,
		num: 0,
		right: 3,
		down: 1,
		y: 0,
	}

	fmt.Println(trees.calcSlope(split))

	slopes := []struct{ x, y int }{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	var res int
	for _, slope := range slopes{
		trees := Trees{
			x: 0,
			num: 0,
			right: slope.x,
			down: slope.y,
			y: 0,
		}

		val := trees.calcSlope(split)
		fmt.Println(val)
		if res == 0 {
			res = val
		}else{
			res *= val
		}
	}
	fmt.Println(res)

}