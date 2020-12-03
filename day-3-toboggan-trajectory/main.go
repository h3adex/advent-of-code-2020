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

func (t *Trees) countTreesInArea(area [] string) int{
	for t.x < len(area)-1 {
		//35 -> #
		if area[t.x][t.y] == '#' {
			t.num++
		}
		t.y = (t.y + t.right) % len(area[t.x])
		t.x += t.down
	}
	return t.num
}


func main() {
	content, err := ioutil.ReadFile("day-3-toboggan-trajectory/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	split := strings.Split(string(content), "\n")

	trees := Trees{
		x: 0,
		num: 0,
		right: 3,
		down: 1,
		y: 0,
	}

	fmt.Println(trees.countTreesInArea(split))

}