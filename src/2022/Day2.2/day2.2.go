package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
2.2 :
Round should end the follwing way:
loss = 0
draw = 3
win  = 6

X = loss
Y = draw
Z = win

A = 1 Rock
B = 2 Paper
C = 3 Scissors

Input file is in the same directory: input.txt  This will provide the game play.
The data consists of two columns, Player1 and Player2.  Player2 is us.
What is the total for player2
*/

const (
	loss = 0
	draw = 1
	win = 2
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
		var Round string

		// Line structure "C X"
		line_str := strings.TrimSpace(fileScanner.Text())
		g_hand := strings.Fields(line_str)

		p1_hand := gethandplayed((g_hand[0]))
		p2_hand := gethandplayed((g_hand[1]))

		game_result := calcplayer(p1_hand, p2_hand)

		var round_score int = 0
		if p2_hand == loss {
			round_score = 0 + game_result
			g.loss += round_score
			Round = "loss"
		}else if p2_hand == draw {
			round_score = 3 + game_result
			g.draw += round_score
			Round = "draw"
		}else if p2_hand == win {
			round_score = 6 + game_result
			g.draw += round_score
			Round = "win"
		}
		fmt.Println("Round is a: ", p2_hand,"- ",Round, " Round Score: ", round_score)
		g.total_score += round_score

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
		"X": 0, // loss
		"Y": 1, // draw
		"Z": 2, // win
	}
	result = game_option[selection]
	return result
}

func calcplayer(p1,p2 int) int {
	/*
	P2:
	X Loss  = 0
	Y Draw  = 3
	Z Win   = 6

	P1:
	Rock    = 1
	Paper   = 2
	Scissor = 3

	Solution: 
	Rock{1} and P2{Y}(Draw) select {1}
	Paper{2} and P2{X}(Loss) select {1}
	Scissors{3} and P2{Z}(Win) select {1}


	Game Rules: 
	 Rock{1} defeats Scissors{3}
	 Scissors{3} defeats Paper{2}
	 Paper{2} defeats Rock{1}.
		          P1
		       |n|R|P|S
			   +-+-+-+-+
	   [X] loss|0|S|R|P|
	           +-+-+-+-+
	P2 [Y] draw|0|R|P|S|
	           +-+-+-+-+
	   [Z] win |0|P|S|R|

	*/
	var p_result int

	// this is based on P2 as the row and P1 as the column
	game_matrix := [][]int{}
	row0 := []int{0,scissors, rock, paper}
	row1 := []int{0,rock, paper, scissors}
	row2 := []int{0,paper,scissors,rock}
	game_matrix = append(game_matrix, row0)
	game_matrix = append(game_matrix, row1)
	game_matrix = append(game_matrix, row2)
	p_result = game_matrix[p2][p1]

	return p_result
}