package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func compareIntSlices(a []int, b []int) bool {
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

func includeIntSlices(a [][]int, b []int) bool {
	for i := 0; i < len(a); i++ {
		if compareIntSlices(a[i], b) {
			return true
		}
	}
	return false
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
	fmt.Println(rules)

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
	fmt.Println(updates)

	count := 0
	for i := 0; i < len(updates); i++ {
		update := updates[i]
		for j := 0; j < len(updates[i])-1; j++ {
			tmp := []int{update[j], update[j+1]}
			if !includeIntSlices(rules, tmp) {
				break
			}
			if j == len(updates[i])-2 {
				count++
			}
		}
	}
	fmt.Println(count)
}
