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
	input := ""
	for scanner.Scan() {
		input += scanner.Text()
	}
	input = "do()" + input + "do()" + "don't()"
	count := 0

	for strings.Index(input, "mul(") != -1 {
		start := strings.Index(input, "mul(")
		mid := start + strings.Index(input[start:], ",")
		end := mid + strings.Index(input[mid:], ")")
		do := strings.Index(input, "do()")
		dont := strings.Index(input, "don't()")

		if dont < start && (do < dont || start < do) {
			input = "don't()" + input[start+1:]
			continue
		}

		firstPart, err := strconv.Atoi(input[start+len("mul(") : mid])
		if err != nil {
			input = input[start+1:]
			println("the former is not a number")
			continue
		} else {
			fmt.Println(firstPart)
		}

		secondPart, err := strconv.Atoi(input[mid+1 : end])

		if err != nil {
			input = input[mid+1:]
			println("the latter is not a number")
			continue
		} else {
			fmt.Println(secondPart)
		}
		count += firstPart * secondPart
		input = input[end+1:]
	}
	fmt.Printf("the total is %v\n", count)
}
