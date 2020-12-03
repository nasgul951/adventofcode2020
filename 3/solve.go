package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		exit("could not open input.txt!")
	}
	defer file.Close()
	s := bufio.NewScanner(file)
	var values []string

	i := 0
	for s.Scan() {
		i++
		line := s.Text()
		values = append(values, line)
	}

	slopes := [5][3]int{
		{1, 1, 0},
		{3, 1, 0},
		{5, 1, 0},
		{7, 1, 0},
		{1, 2, 0},
	}

	for i, s := range slopes {
		fmt.Printf("Runing slope %d", i)
		s[2] = runSlope(values, s[0], s[1])
		slopes[i] = s
		fmt.Printf("Finished slope, hit %d trees\n", s[2])
	}

	baseProduct := 1
	for i, s := range slopes {
		fmt.Printf("%d: Finished slope right %d, down %d - encountered %d trees\n", i, s[0], s[1], s[2])
		baseProduct *= s[2]
	}

	fmt.Printf("Product of trees is: %d\n", baseProduct)
}

func runSlope(values []string, rt int, dn int) int {
	x, y, treeCount := 0, 0, 0
	for y < len(values) {
		if printSlope(values[y], x) {
			treeCount++
		}

		x += rt
		y += dn
	}

	return treeCount
}

func printSlope(s string, x int) bool {
	i := 0
	for i < (x / 31) {
		fmt.Printf("%s", s)
		i++
	}

	ix := x % 31
	var s1 string
	var isTree bool
	if s[ix] == "#"[0] {
		s1 = replaceByteInString(s, ix, "X"[0])
		isTree = true
	} else {
		s1 = replaceByteInString(s, ix, "O"[0])
		isTree = false
	}
	fmt.Printf("%s\n", s1)

	return isTree
}

func replaceByteInString(s string, ix int, b byte) string {
	var s1 string
	bytes := []byte(s)
	bytes[ix] = b
	s1 = string(bytes)

	return s1
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
