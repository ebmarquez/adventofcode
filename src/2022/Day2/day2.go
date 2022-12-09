/*
Day 2
Find the top three Elves carrying the most Calories. How many Calories are those Elves carrying in total?
*/

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

	// this looks like a linked list, creating a stack structure
	// Last in first out.  this will allow me to capture the top 3 leaders.
	// If the current elf has more calroies than the current leader,
	//    push into the stack and repeat. 
	// If the current elf does not have more calories then ignore the elf.
	// pop the first three elfs from the linked list and sum them.  This is the answer.

	elf_leader_sum := 0 // sum of leader1
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
		cal += item  // sum the elf's total calories

		// identify if this is a new elf, new elf is zero
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