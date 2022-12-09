/*
Day 2
Find the top three Elves carrying the most Calories. How many Calories are those Elves carrying in total?
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	i := 0     // counter to determine the elf
	total := 0 // calories the elf is holding, this is rest for each elf
	var elfTotal = []int{}
	slc_total := append(elfTotal, 0)

	readFile, err := os.Open("testdata.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line_str := strings.TrimSpace(fileScanner.Text())
		item, _ := strconv.Atoi(line_str)

		// identify if this is a new elf, new elf is zero
		if item == 0 {
			//fmt.Println("Total:", total)
			slc_total = append(slc_total, total)
			//fmt.Println(slc_total)
			i++       // increment to the next elf
			total = 0 // new elf, rest back to zero
		} else {
			total += item
		}
	}
	readFile.Close()

	sort.Ints(slc_total)
	//fmt.Println("length: ", l)
	l := len(slc_total) - 1
	end := l - 2
	var total_cal int
	for ; l >= end; l--  {
		fmt.Println(slc_total[l])
		total_cal += slc_total[l]
	}
	fmt.Println(slc_total)
	fmt.Println("total: ", total_cal)
}