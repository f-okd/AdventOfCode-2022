package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ElvesWithMostCalories(elf_list map[int]int) int {
	topElves := make([]int, 3)
	for _, calories := range elf_list {
		if calories > topElves[0] {
			topElves[2] = topElves[1]
			topElves[1] = topElves[0]
			topElves[0] = calories
		} else if calories > topElves[1] {
			topElves[2] = topElves[1]
			topElves[1] = calories
		} else if calories > topElves[2] {
			topElves[2] = calories
		}
	}
	sum := topElves[0] + topElves[1] + topElves[2]
	return sum
}

func main() {
	file, err := os.Open("Input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	elf_list := make(map[int]int)
	var total int

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if total != 0 {
				elf_list[len(elf_list)+1] = total
				total = 0
			}
			continue
		}

		num, err := strconv.Atoi(strings.TrimSpace(line))
		if err != nil {
			log.Fatal(err)
		}

		total += num
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// check if remaining total unaccounted for (last line, varies if there's a slash n at the end)
	if total != 0 {
		elf_list[len(elf_list)+1] = total
	}

	for id, calories := range elf_list {
		fmt.Printf("Elf %d carrying food with %d calories\n", id, calories)
	}

	calories_carried := ElvesWithMostCalories(elf_list)
	fmt.Println("The top 3 elves are carrying", calories_carried, " calories:")
}
