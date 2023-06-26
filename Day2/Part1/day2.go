package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

/*
A = Opponent rock
B = Opponent paper
c = Opponent scissors

X = My rock = + 1 shape score
Y = My paper = + 2 shape score
Z = My scissors = + 3 shape score

Points per match:
0 for loss
3 for tie
6 for win
*/
func mactch(OPPONENT, YOU byte) int {
	switch {
	case OPPONENT == 'A' && YOU == 'X':
		return 1 + 3
	case OPPONENT == 'A' && YOU == 'Y':
		return 2 + 6
	case OPPONENT == 'A' && YOU == 'Z':
		return 3 + 0
	case OPPONENT == 'B' && YOU == 'X':
		return 1 + 0
	case OPPONENT == 'B' && YOU == 'Y':
		return 2 + 3
	case OPPONENT == 'B' && YOU == 'Z':
		return 3 + 6
	case OPPONENT == 'C' && YOU == 'X':
		return 1 + 6
	case OPPONENT == 'C' && YOU == 'Y':
		return 2 + 0
	case OPPONENT == 'C' && YOU == 'Z':

		return 3 + 3
	default:
		return 0
	}
}

func main() {
	file, err := os.Open("Input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	SCORE := 0

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) > 2 {
			OPPONENT := line[0]
			YOU := line[2]
			SCORE += mactch(OPPONENT, YOU)
		}
	}

	fmt.Println(SCORE)
}
