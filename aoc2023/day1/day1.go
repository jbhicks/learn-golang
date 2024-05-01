package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	var digits []int
	file, err := os.Open("input")
	if err != nil {
		fmt.Println("Error reading file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		currentLine := scanner.Text()
		var temp string

		for _, char := range currentLine {
			if unicode.IsDigit(char) {
				temp += string(char)
			}
		}

		// take only first and last digits
		if len(temp) > 0 {
			temp = temp[:1] + temp[len(temp)-1:]
		}

		// Convert string to int
		num, err := strconv.Atoi(temp)
		if err != nil {
			fmt.Println("Error converting string to int")
		}

		digits = append(digits, num)
		fmt.Println(digits)
		temp = ""
	}

	// Sum of all digits
	sum := 0
	for _, digit := range digits {
		sum += digit
	}
	fmt.Println("Sum of all digits: ", sum)
}
