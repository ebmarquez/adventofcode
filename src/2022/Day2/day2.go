package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Highlights:
A = X = 1 // Rock
B = Y = 2 // Paper
C = Z = 3 // Sissor

Loss = 0
Draw = 3
Win  = 6

Input file is in the same directory: input.txt  This will provide the game play.
The data consists of two columns, Player1 and Player2.  Player2 is us.
What is the total for player2
*/

const (
	loss = 0
	draw = 3
	win = 6
	rock = 1
	paper = 2
	scissors = 3
)

// Defining a struct type
type game struct {
	loss, draw, win, total_score, total_played int
}

func main() {

	var g game
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {

		// Line structure "C X"
		line_str := strings.TrimSpace(fileScanner.Text())
		g_hand := strings.Fields(line_str)

		p1_hand := gethandplayed((g_hand[0]))
		p2_hand := gethandplayed((g_hand[1]))

		game_result := calcplayer(p1_hand, p2_hand)

		if game_result == loss {
			g.loss += loss + p2_hand
		}else if game_result == draw {
			g.draw += draw + p2_hand
		}else if game_result == win {
			g.win += win + p2_hand
		}
		//var round_score int
		var round_score int = 0
		round_score = (p2_hand + game_result)
		g.total_score += round_score
		fmt.Printf("(p2_hand: %d + result: %d) = roundscore(%d) - Total Score: [%d]:%d\n", p2_hand, game_result, round_score, g.total_played,g.total_score)

		g.total_played++
		
	}
	readFile.Close()
	fmt.Println(g.total_score)
	total_result := g.draw + g.loss + g.win
	fmt.Printf("Total_result (%d)\n", total_result)
}

func gethandplayed(selection string) int {
	/*
		This will return a interger matching the game selection.
	*/
	var result int
	game_option := map[string]int{
		"A": rock,
		"B": paper,
		"C": scissors,
		"X": rock,
		"Y": paper,
		"Z": scissors,
	}
	result = game_option[selection]
	return result
}

func calcplayer(p1,p2 int) int {
	/*
	Loss = 0
	Draw = 3
	Win  = 6

	Rules: 
	 Rock{1} defeats Scissors{3}
	 Scissors{3} defeats Paper{2}
	 Paper{2} defeats Rock{1}.
		   P1
	    |x|R|P|S
	   x|0|0|0|0|
	P2 R|0|3|0|6|
	   P|0|6|3|0|
	   S|0|0|6|3|
	*/
	var p_result int

	// this is based on P2 as the row and P1 as the column
	game_matrix := [][]int{}
	row0 := []int{0,0,0,0}  // not needed
	row1 := []int{0,draw, loss, win}
	row2 := []int{0,win,draw,loss}
	row3 := []int{0,loss,win,draw}
	game_matrix = append(game_matrix, row0)
	game_matrix = append(game_matrix, row1)
	game_matrix = append(game_matrix, row2)
	game_matrix = append(game_matrix, row3)
	p_result = game_matrix[p2][p1]

	return p_result
}