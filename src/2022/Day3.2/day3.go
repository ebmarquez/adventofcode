package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	//"unicode"
)

/*
The first rucksack contains the items `vJrwpWtwJgWrhcsFMMfFFhFp`, which means its first compartment contains the items `vJrwpWtwJgWr`, while the second compartment contains the items `hcsFMMfFFhFp`. The only item type that appears in both compartments is lowercase p.

- divide the line in half.
- compaire the two halves and look for a repeat char
- identify the letter priority
	- Lowercase item types a through z have priorities 1 through 26.
	- Uppercase item types A through Z have priorities 27 through 52.
- Return the sum of all the priority numbers

*/

func main() {
	// Open the file
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	// read the line
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var sum int
	for fileScanner.Scan() {
		line_str := strings.TrimSpace(fileScanner.Text())
		line_len := len(line_str)
		// strings.Split(line_str,line_len)
		//fmt.Printf("line len: [%d], half [%d]\n", line_len, line_len/2)

		line := split_string(line_len/2, line_str)
		//fmt.Println(line)
		item := test(line)
		priority := assignValueOfString(item)
		//fmt.Println("Value: ", priority)


		sum += priority
		fmt.Printf("(%02d) sum: %05d\n", priority, sum,)

	}
}

func split_string(l int, s string) [2]string {
	var my_str [2]string
	str_len := len(s)

	// substring at the half way mark
	my_str[0] = s[0:l]
	my_str[1] = s[l:str_len]

	return my_str
}

func test(str1 [2]string) string {

	var res1 string
	//index := strings.IndexAny(str1[0], "G")
	for _, v := range str1[0] {
		//.Print("i is: ", i)
		//fmt.Print(" - v is: ", string(v))
		index := strings.IndexAny(str1[1], string(v))
		//fmt.Println(" - index of the letter is: ", index)
		if index >= 0 {
			res1 = string(v)
			break
		}
	}
	return res1
}

func assignValueOfString(str string) int {
	// var result int
	// fmt.Println(str)
	// ascii_val := byte(str)

	// if strings.ToUpper(str) == str {
	// 	pri := int(byte(ascii_val)) - 38
	// } else {
	// 	// must be lowercase
	// }

	//response := unicode.IsUpper(chars)
	//fmt.Println("isUpper: ", test)
	letter_map := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
		"g": 7,
		"h": 8,
		"i": 9,
		"j": 10,
		"k": 11,
		"l": 12,
		"m": 13,
		"n": 14,
		"o": 15,
		"p": 16,
		"q": 17,
		"r": 18,
		"s": 19,
		"t": 20,
		"u": 21,
		"v": 22,
		"w": 23,
		"x": 24,
		"y": 25,
		"z": 26,
		"A": 27,
		"B": 28,
		"C": 29,
		"D": 30,
		"E": 31,
		"F": 32,
		"G": 33,
		"H": 34,
		"I": 35,
		"J": 36,
		"K": 37,
		"L": 38,
		"M": 39,
		"N": 40,
		"O": 41,
		"P": 42,
		"Q": 43,
		"R": 44,
		"S": 45,
		"T": 46,
		"U": 47,
		"V": 48,
		"W": 49,
		"X": 50,
		"Y": 51,
		"Z": 52,
	}
	result := letter_map[str]
	//fmt.Println("letter value: ", result)
	return result
}
