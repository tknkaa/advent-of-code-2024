package main

import (
	"fmt"
	"strconv"
)

func findFirstBlock(num int, files []int) (int, int) {
	start := 0
	count := 0
	for i := 0; i < len(files); i++ {
		if files[i] == num {
			start = i
			break
		}
	}
	for i := 0; i < len(files); i++ {
		if files[i] == num {
			count++
		} else if count != 0 && files[i] != num {
			break
		}
	}
	return start, count
}

func findEmptyBlocks(files []int) [][]int {
	emptyBlocks := make([][]int, 0)
	i := 0
	idx, count := findFirstBlock(-1, files)
	for count > 0 {
		emptyBlocks = append(emptyBlocks, []int{idx, count})
		i = idx + count
		if i >= len(files) {
			break
		}
		nextIdx, nextCount := findFirstBlock(-1, files[i:])
		idx = nextIdx + i
		count = nextCount
	}
	return emptyBlocks
}

func findPlace(files []int, num int) (int, bool) {
	idx, minCount := findFirstBlock(num, files)
	emptyBlocks := findEmptyBlocks(files[:idx])
	for _, block := range emptyBlocks {
		if block[1] >= minCount {
			return block[0], true
		}
	}
	return -1, false
}
func replaceNum(files []int, num int) {
	idx, count := findFirstBlock(num, files)
	nextIdx, canReplace := findPlace(files, num)
	if !canReplace {
		return
	}
	for i := nextIdx; i < nextIdx+count; i++ {
		files[i] = num
	}
	for i := idx; i < idx+count; i++ {
		files[i] = -1
	}
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
	max := (len(nums) - 1) / 2
	for n := max; n > 0; n-- {
		replaceNum(files, n)
	}
	count := 0
	for i, v := range files {
		if v != -1 {
			count += i * v
		}
	}
	fmt.Println(count)
}
