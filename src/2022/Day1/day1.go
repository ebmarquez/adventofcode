package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)
/*
* Sample Data
1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
*/
func main() {

	elf_leader_sum := 0 // sum of the leader
	elf_leader := 0		// elf leader number
	elf := 0			// counter to determine the elf
	cal := 0 			// calories the elf is holding, this is rest for each elf
	
	readFile, err := os.Open("testdata.txt")
    if err != nil {
        fmt.Println(err)
    }

    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)

    for fileScanner.Scan() {
        line_str := strings.TrimSpace(fileScanner.Text())
		item, _ := strconv.Atoi(line_str)
		cal += item

		if item == 0 {
			fmt.Printf("Elf [%d] had [%d] Calories\n",elf,cal)
			if cal > elf_leader_sum {
				fmt.Printf(" -- New Leader Elf [%d] --Previous Elf [%d] had [%d] Calories\n", elf,elf_leader,elf_leader_sum)
				elf_leader_sum = cal
				elf_leader = elf
			}
			elf++	// increment to the next elf
			cal = 0	// new elf, rest back to zero
		}
    }
	fmt.Printf("\nLead EFl [%d] had [%d] calories\n",elf_leader, elf_leader_sum)
  
    readFile.Close()

}