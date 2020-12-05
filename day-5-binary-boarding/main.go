package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main(){
	content, err := ioutil.ReadFile("day-5-binary-boarding/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	split := strings.Split(strings.TrimSpace(string(content)), "\n")

	var maxSeat int64
	ids := make([]int, 0)

	for _, str := range split {
		r := strings.NewReplacer("F", "0", "B", "1", "L", "0", "R", "1")
		str = r.Replace(str)
		row, _ := strconv.ParseInt(str[:7], 2, 64)
		col, _ := strconv.ParseInt(str[7:], 2, 64)

		ids = append(ids, int(row*8+col))

		if (row*8+col) > maxSeat {
			maxSeat = row*8+col
		}
	}

	fmt.Printf("Max seats %d \n", maxSeat)
	sort.Ints(ids)

	for i := 0; i < len(ids)-1; i++ {
		if ids[i+1] == ids[i]+2 {
			fmt.Printf("My seat %d \n", ids[i]+1)
		}
	}
}
