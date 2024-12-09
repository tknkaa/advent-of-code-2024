package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type Vertex struct {
	x int
	y int
}

func findAntinode(width int, height int, loc1 Vertex, loc2 Vertex) []Vertex {
	dy := loc2.y - loc1.y
	dx := loc2.x - loc1.x
	antinodes := make([]Vertex, 0)
	y1 := loc1.y - dy
	x1 := loc1.x - dx
	y2 := loc2.y + dy
	x2 := loc2.x + dx

	for 0 <= y1 && y1 < height && 0 <= x1 && x1 < width {
		antinodes = append(antinodes, Vertex{y: y1, x: x1})
		y1 -= dy
		x1 -= dx
	}
	for 0 <= y2 && y2 < height && 0 <= x2 && x2 < width {
		antinodes = append(antinodes, Vertex{y: y2, x: x2})
		y2 += dy
		x2 += dx
	}
	return antinodes
}

func chooseAntennaPairs(antennas []Vertex) [][]Vertex {
	antennaPairs := make([][]Vertex, 0)
	for i := 0; i < len(antennas)-1; i++ {
		for j := i + 1; j < len(antennas); j++ {
			antennaPairs = append(antennaPairs, []Vertex{antennas[i], antennas[j]})
		}
	}
	return antennaPairs
}

func findAntennas(antenna rune, field [][]rune) []Vertex {
	if len(field) == 0 {
		return []Vertex{}
	}
	locations := make([]Vertex, 0)
	for y := 0; y < len(field); y++ {
		for x := 0; x < len(field[0]); x++ {
			if field[y][x] == antenna {
				locations = append(locations, Vertex{y, x})
			}
		}
	}
	return locations
}

func judgeContains(prev []Vertex, new Vertex) bool {
	for _, one := range prev {
		if one.y == new.y && one.x == new.x {
			return true
		}
	}
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	field := make([][]rune, 0)
	for scanner.Scan() {
		field = append(field, []rune(scanner.Text()))
	}
	emptyFlag := []rune(".")[0]
	width := len(field[0])
	height := len(field)

	antennas := make([]rune, 0)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if field[y][x] != emptyFlag && !slices.Contains(antennas, field[y][x]) {
				antennas = append(antennas, field[y][x])
			}
		}
	}

	antinodes := make([]Vertex, 0)

	//同じコード繰り返していてあんまよくないけど、後からアンテナの位置をappendすると重複しそう
	for _, anntenna := range antennas {
		locations := findAntennas(anntenna, field)
		antinodes = append(antinodes, locations...)
	}

	for _, antenna := range antennas {
		locations := findAntennas(antenna, field)
		pairs := chooseAntennaPairs(locations)
		for _, pair := range pairs {
			antinodePairs := findAntinode(width, height, pair[0], pair[1])
			for i := 0; i < len(antinodePairs); i++ {
				if !judgeContains(antinodes, antinodePairs[i]) {
					antinodes = append(antinodes, antinodePairs[i])
				}
			}
			fmt.Println("antenna", antenna, "pair", pair, "antinodePair", antinodePairs)
		}
	}
	fmt.Println("antinode", antinodes)
	fmt.Println("count", len(antinodes))
}
