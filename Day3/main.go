package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	count := 0

	for strings.Index(input, "mul(") != -1 {
		start := strings.Index(input, "mul(")
		mid := start + strings.Index(input[start:], ",")
		end := mid + strings.Index(input[mid:], ")")

		firstPart, err := strconv.Atoi(input[start+len("mul(") : mid])
		if err != nil {
			input = input[end+1:]
			println("the former is not a number")
			continue
		} else {
			fmt.Println(firstPart)
		}

		secondPart, err := strconv.Atoi(input[mid+1 : end])

		if err != nil {
			input = input[end+1:]
			println("the latter is not a number")
			continue
		} else {
			fmt.Println(secondPart)
		}
		count += firstPart + secondPart
		input = input[end+1:]
	}
	fmt.Printf("the total is %v\n", count)
}
