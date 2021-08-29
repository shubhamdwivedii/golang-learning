package main

import (
	"fmt"
	"regexp"
)

func main() {

	re, _ := regexp.Compile("shubham")

	str := "My name is Shubham"

	match := re.FindStringIndex(str)
	fmt.Println(match) // [] // Because Shubham starts with Capital S

	str2 := "I am shubham"
	match2 := re.FindStringIndex(str2)
	fmt.Println(match2) // [5 12] // from (including) - Upto (exclution)

	match3 := re.FindString(str2)
	fmt.Println(match3) // shubham

	str3 := "shubham is my name, yes shubham"
	match4 := re.FindAllStringSubmatchIndex(str3, -1)
	fmt.Println(match4) // [[0 7] [24 31]]

	re2, _ := regexp.Compile("[0-9]+-s.*d")
	match5 := re2.FindString("20024-shubham_dubey")
	fmt.Println(match5) // 20024-shubham_d
}
