package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func maxCalories(elf map[int]int) (int, int) {
	max := -1
	fatass_elf := -1
	for elf_id, calories := range elf {
		if calories > max {
			max = calories
			fatass_elf = elf_id
		}
	}
	return fatass_elf, max
}

func main() {
	file, err := os.Open("Input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	elf := make(map[int]int)
	var total int

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if total != 0 {
				elf[len(elf)+1] = total
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
		elf[len(elf)+1] = total
	}

	for id, calories := range elf {
		fmt.Printf("Elf %d carrying food with %d calories\n", id, calories)
	}

	elf_id, calories_carried := maxCalories(elf)
	fmt.Println("Elf", elf_id, "is carrying the most calories:", calories_carried)
}
