package main

import (
	"fmt"
	"strconv"
)

func findFirstEmpty(files []int) int {
	for i := 0; i < len(files); i++ {
		if files[i] == -1 {
			return i
		}
	}
	return -1 //含まれない
}
func main() {
	var input string
	fmt.Scanf("%s", &input)
	runes := []rune(input)
	nums := make([]int, 0)
	for _, v := range runes {
		num, _ := strconv.Atoi(string(v))
		nums = append(nums, num)
	}
	files := make([]int, 0)
	for i, v := range nums {
		if i%2 == 0 {
			for j := 0; j < v; j++ {
				files = append(files, i/2)
			}
		} else {
			for j := 0; j < v; j++ {
				files = append(files, -1)
			}
		}
	}
	idx := findFirstEmpty(files)
	for idx != -1 {
		files[idx] = files[len(files)-1]
		files = files[:len(files)-1]
		idx = findFirstEmpty(files)
	}
	count := 0
	for i, v := range files {
		count += i * v
	}
	fmt.Println(count)
}
