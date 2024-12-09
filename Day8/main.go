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
	if 0 <= loc1.y-dy && loc1.y-dy < height && 0 <= loc1.x-dx && loc1.x-dx < width {
		antinodes = append(antinodes, Vertex{y: loc1.y - dy, x: loc1.x - dx})
	}
	if 0 <= loc2.y+dy && loc2.y+dy < height && 0 <= loc2.x+dx && loc2.x+dx < width {
		antinodes = append(antinodes, Vertex{y: loc2.y + dy, x: loc2.x + dx})
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

func findAntennas(antenna byte, field [][]byte) []Vertex {
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
	field := make([][]byte, 0)
	for scanner.Scan() {
		field = append(field, scanner.Bytes())
	}
	emptyFlag := []byte(".")[0]
	width := len(field[0])
	height := len(field)

	antennas := make([]byte, 0)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if field[y][x] != emptyFlag && !slices.Contains(antennas, field[y][x]) {
				antennas = append(antennas, field[y][x])
			}
		}
	}

	antinodes := make([]Vertex, 0)
	for _, antenna := range antennas {
		locations := findAntennas(antenna, field)
		pairs := chooseAntennaPairs(locations)
		for _, pair := range pairs {
			antinodePairs := findAntinode(width, height, pair[0], pair[1])
			if len(antinodePairs) == 0 {
				continue
			} else if len(antinodePairs) == 1 && !judgeContains(antinodes, antinodePairs[0]) {
				antinodes = append(antinodes, antinodePairs[0])
			} else if len(antinodePairs) == 2 {
				if !judgeContains(antinodes, antinodePairs[0]) {
					antinodes = append(antinodes, antinodePairs[0])
				}
				if !judgeContains(antinodes, antinodePairs[1]) {
					antinodes = append(antinodes, antinodePairs[1])
				}
			} else {
				continue
			}
			fmt.Println("antenna", antenna, "pair", pair, "antinodePair", antinodePairs)
		}
	}
	fmt.Println("antinode", antinodes)
	fmt.Println("count", len(antinodes))
}
