package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

/*
The total score is still calculated in the same way, but now you need to figure out what shape to choose so the round ends as indicated
The second column says how the round needs to end:
- X means you need to lose
- Y means you need to end the round in a draw
- Z means you need to win.
*/
// // Returns what you should pick against the opponent to get the desired outcome
func strategy(OPPONENT, DESIRED_OUTCOME byte) byte {
	BEAT := make(map[byte]byte)
	BEAT['A'] = 'Y'
	BEAT['B'] = 'Z'
	BEAT['C'] = 'X'

	DRAW_AGAINST := make(map[byte]byte)
	DRAW_AGAINST['A'] = 'X'
	DRAW_AGAINST['B'] = 'Y'
	DRAW_AGAINST['C'] = 'Z'

	LOSE_TO := make(map[byte]byte)
	LOSE_TO['A'] = 'Z'
	LOSE_TO['B'] = 'X'
	LOSE_TO['C'] = 'Y'

	switch {
	case DESIRED_OUTCOME == 'X':
		return LOSE_TO[OPPONENT]
	case DESIRED_OUTCOME == 'Y':
		return DRAW_AGAINST[OPPONENT]
	case DESIRED_OUTCOME == 'Z':
		return BEAT[OPPONENT]
	default:
		return 'X'
	}
}

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
// Returns the score for a match
func match(OPPONENT, YOU byte) int {
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
			DESIRED_OUTCOME := line[2]
			YOU := strategy(OPPONENT, DESIRED_OUTCOME)
			SCORE += match(OPPONENT, YOU)
		}
	}

	fmt.Println(SCORE)
}
