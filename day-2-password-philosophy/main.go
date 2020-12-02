package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Password struct {
	number1 int
	number2 int
	letter string
	password string
}

func (p *Password) policyV1() (int, error){
	var counter = 0
	for _, char := range p.password {
		if string(char) == p.letter {
			counter ++
		}
	}
	return counter, nil
}

func (p *Password) parsePasswordString(s string){
	fields := strings.FieldsFunc(s, func(r rune) bool {
		return strings.ContainsRune(" :-", r)
	})
	p.number1, _ = strconv.Atoi(fields[0])
	p.number2, _ = strconv.Atoi(fields[1])
	p.letter = fields[2]
	p.password = fields[3]

}

func (p *Password) policyV2() error{
	l1 := string(p.password[p.number1-1]) == p.letter
	l2 := string(p.password[p.number2-1]) == p.letter
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
	var policySum2 = 0
	var password = Password{}

	for _, line := range input {
		password.parsePasswordString(line)
		num, err := password.policyV1()
		if err != nil {}

		if num >= password.number1 && num <= password.number2 {
			policySum1++
		}

		if password.policyV2() == nil{
			policySum2++
		}
	}
	fmt.Printf("Policy 1: found %d policy1 \n", policySum1)
	fmt.Printf("Policy 2: found %d policy2 \n", policySum2)
}