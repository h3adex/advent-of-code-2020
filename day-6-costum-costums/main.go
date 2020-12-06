package main

import (
	"fmt"
	"go/types"
	"io/ioutil"
	"strings"
)


func main(){
	content, err := ioutil.ReadFile("day-6-costum-costums/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	//split \n\n
	split := strings.Split(strings.TrimSpace(string(content)), "\n\n")

	var counter = 0
	for _, curr := range split {
		//new map for each group
		var Answers = make(map[rune] types.Nil)
		for _, line := range strings.Split(curr, "\n") {
			for _, char := range line {
				//exists -> true or false
				if _, exists := Answers[char]; !exists {
					Answers[char] = types.Nil{}
					counter++
				}
			}
		}
	}
	fmt.Printf("Answer 1: %d \n", counter)

	var counter2 = 0
	for _, curr := range split {
		//collect all given answers
		Answers := make(map[rune]int)
		lines := strings.Split(curr, "\n")

		for _, line := range lines {
			for _, char := range line {
				Answers[char]++
			}
		}

		count := 0
		//Example: su \n egu
		//map: 115(s) -> 1 117(u) -> 2  101(e) -> 1 103(u) -> 1
		for chars := range Answers {

			//Example k:1 v:true
			//k, v := givenAnswers[question]
			//fmt.Println(k, v)

			//check if answer exits in map && check if everyone answered this question with yes
			if ansCount, exists := Answers[chars]; exists && (ansCount == len(lines)) {
				count++
			}
		}
		counter2 += count
	}

	fmt.Printf("Answer 2: %d \n", counter2)
}