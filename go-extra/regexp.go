package main

import (
	"fmt"
	"regexp"
)

func main() {
	name := "shubhamshubham"

	// Returns true if pattern is present
	same, err := regexp.MatchString("shubham", name)
	fmt.Println(same, err) // true nil

	re, _ := regexp.Compile("shubham")

	str := "My name is Shubham"

	// Returns slice of first and last index
	match := re.FindStringIndex(str)
	fmt.Println(match) // [] // Because Shubham starts with Capital S

	str2 := "I am shubham"
	match2 := re.FindStringIndex(str2)
	fmt.Println(match2) // [5 12] // from (including) - Upto (exclution)

	// returns matching string
	match3 := re.FindString(str2)
	fmt.Println(match3) // shubham

	str3 := "shubham is my name, yes shubham"
	// returns a slice of slices of first and last index of every matching string.
	match4 := re.FindAllStringSubmatchIndex(str3, -1)
	fmt.Println(match4) // [[0 7] [24 31]]

	re2, _ := regexp.Compile("[0-9]+-s.*d")
	match5 := re2.FindString("20024-shubham_dubey")
	fmt.Println(match5) // 20024-shubham_d

	re3, _ := regexp.Compile(" ")
	// Will replace ever " " with "+"
	match6 := re3.ReplaceAllString("I am very good in Golang.", "+")
	fmt.Println(match6) // I+am+very+good+in+Golang.

	validateIp("200.106.141.15")
	validateIp("160.103.7.140")
	validateIp("70.95.73.73")
	validateIp("1.0.01.34") // invalid due to 01
}

func validateIp(ip string) bool {
	re, _ := regexp.Compile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`)
	res := re.MatchString(ip)
	fmt.Println(res)
	return res
}
