package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type Passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

var r = regexp.MustCompile(`(eyr|hgt|pid|ecl|cid|hcl|byr|iyr):(.*?)\s|\n`)

func (p *Passport) policy1() bool{
	//ignore cid
	return p.byr != "" &&
			p.iyr != "" &&
			p.eyr != "" &&
			p.hgt != "" &&
			p.hcl != "" &&
			p.ecl != "" &&
			p.pid != ""
}

func (p *Passport) policy2() bool{
	byrVal, _ := strconv.Atoi(p.byr)
	if !(1920 <= byrVal && byrVal <= 2002){
		return false
	}

	iyrVal, _ := strconv.Atoi(p.iyr)
	if !(2010 <= iyrVal && iyrVal <= 2020){
		return false
	}

	eyrVal, _ := strconv.Atoi(p.eyr)
	if !(2020 <= eyrVal && eyrVal <= 2030){
		return false
	}

	var re = regexp.MustCompile("([0-9]+)([a-z]+)")
	matches := re.FindStringSubmatch(p.hgt)
	var hgtValid bool
	if len(matches) >= 2 {
		if matches[2] == "cm" || matches[2] == "in" {
			value, _ := strconv.Atoi(matches[1])
			hgtValid = (matches[2] == "cm" && value >= 150 && value <= 193) || (matches[2] == "in" && value >= 59 && value <= 76)
		}
	}
	hclValid, _ := regexp.MatchString("^#[0-9a-f]{6}$", p.hcl)
	eclValid, _ := regexp.MatchString("^amb|blu|brn|gry|grn|hzl|oth", p.ecl)
	pidValid, _ := regexp.MatchString("^[0-9]{9}$", p.pid)
	if !(pidValid && eclValid && hclValid && hgtValid){
		return false
	}

	return true

}

func main(){
	content, err := ioutil.ReadFile("day-4-passport-processing/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	split := strings.Split(strings.TrimSpace(string(content)), "\n\n")
	passports := make([]Passport, 0)

	for _, v := range split {
		//Add \n to catch EOF
		groups := r.FindAllStringSubmatch(v+"\n", -1)
		fieldMap := make(map[string]string, 0)
		for _, group := range groups{
			fieldMap[group[1]] = group[2]
		}
		passports = append(passports, Passport{
			byr: fieldMap["byr"],
			iyr: fieldMap["iyr"],
			eyr: fieldMap["eyr"],
			hgt: fieldMap["hgt"],
			hcl: fieldMap["hcl"],
			ecl: fieldMap["ecl"],
			pid: fieldMap["pid"],
			cid: fieldMap["cid"],
		})
	}
	result1 := 0
	result2 := 0
	for _, p := range passports {
		if p.policy1(){
			result1++
		}
		if p.policy2(){
			result2++
		}
	}

	fmt.Printf("Result for policy 1: %d \n", result1)
	fmt.Printf("Result for policy 2: %d \n", result2)
}
