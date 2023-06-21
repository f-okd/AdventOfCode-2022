package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("Input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	elf := make(map[int]int)
	var total int
	fmt.Println("Enter the elf food list")

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if total != 0 {
				elf[len(elf)+1] = total
			}
			continue
		}

	}

}
