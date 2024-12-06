package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func createReversePairs(update []int) [][]int {
	ans := make([][]int, 0)
	for i := 0; i < len(update)-1; i++ {
		for j := i + 1; j < len(update); j++ {
			ans = append(ans, []int{update[j], update[i]})
		}
	}
	return ans
}

func judgeViolation(rules [][]int, pairs [][]int) bool {
	for i := 0; i < len(pairs); i++ {
		for j := 0; j < len(rules); j++ {
			if compareSlice(pairs[i], rules[j]) {
				return true
			}
		}
	}
	return false
}

func compareSlice(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func findMiddle(update []int) int {
	return update[(len(update)+1)/2-1]
}

func main() {
	rules := make([][]int, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		if input == "" {
			break
		}
		mid := strings.Index(input, "|")
		tmp := make([]int, 2)
		tmp[0], _ = strconv.Atoi(input[:mid])
		tmp[1], _ = strconv.Atoi(input[mid+1:])
		rules = append(rules, tmp)
	}

	updates := make([][]int, 0)
	for scanner.Scan() {
		input := scanner.Text()
		tmp := strings.Split(input, ",")
		nums := make([]int, 0)
		for i := 0; i < len(tmp); i++ {
			num, _ := strconv.Atoi(tmp[i])
			nums = append(nums, num)
		}
		updates = append(updates, nums)
	}

	count := 0

	for i := 0; i < len(updates); i++ {
		update := updates[i]
		if judgeViolation(rules, createReversePairs(update)) {
			continue
		}
		fmt.Println(update)
		fmt.Println(findMiddle(update))
		count += findMiddle(update)
	}
	fmt.Println(count)
}
