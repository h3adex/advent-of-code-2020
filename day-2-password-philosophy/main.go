package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func policyV1(letter string, words string) (int, error){
	var counter = 0
	for _, char := range words {
		if string(char) == letter {
			counter ++
		}
	}
	return counter, nil
}

func MinMax(words string) (int, int, error){
	min, err := strconv.Atoi(strings.Split(words, "-")[0])
	if err != nil{
		return 0, 0, err
	}
	max, err := strconv.Atoi(strings.Split(words, "-")[1])
	if err != nil{
		return 0, 0, err
	}
	return min, max, nil
}

func policyV2(position1 int, position2 int, letter string, words string) error{

	l1 := string(words[position1-1]) == letter
	l2 := string(words[position2-1]) == letter
	if l1 && !l2 || l2 && !l1 {
		return nil
	}
	return fmt.Errorf("new policy doesn't match")
}

func main() {
	file, err := os.Open("day-2-password-philosophy/input.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var input []string

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	_ = file.Close()

	var policySum1 = 0
	for _, line := range input {
		splitString := strings.Split(line, " ")
		num, err := policyV1(strings.Split(splitString[1], ":")[0], splitString[2])
		if err != nil {
			fmt.Print(err)
		}
		min, max, err := MinMax(splitString[0])
		if err != nil {
			fmt.Println(err)
		}
		if num >= min && num <= max {
			policySum1++
		}

	}
	fmt.Printf("Policy 1: found %d policy1 \n", policySum1)


	var policy2 = 0
	for _, line := range input {
		splitString := strings.Split(line, " ")
		pos1, pos2, err := MinMax(splitString[0])
		if err != nil {
			fmt.Println(err)
		}
		if policyV2(pos1, pos2, strings.Split(splitString[1], ":")[0], splitString[2]) == nil{
			policy2++
		}
	}
	fmt.Printf("Policy 2: found %d policy2 \n", policy2)
}