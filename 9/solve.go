package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	numbers := loadData("input.txt")

	invalid := 0
	for i := 25; i < len(numbers); i++ {
		if !numbers.IsValid(i) {
			invalid = numbers[i]
			fmt.Printf("%d at index %d is not valid\n", invalid, i)
		}
	}

	set := numbers.FindSet(invalid)
	if set == nil {
		exit("Failed to find a valid set!")
	}

	fmt.Printf("Found Set:\n%v\n", set)
	min, max := set.Min(), set.Max()
	fmt.Printf("Min: %d, Max: %d, Sum: %d\n", min, max, min+max)
}

type numList []int

func (nl numList) IsValid(ix int) bool {
	n := nl[ix]
	last25 := nl[ix-25 : ix]
	for i, a := range last25 {
		for j, b := range last25 {
			if i != j && a+b == n {
				return true
			}
		}
	}

	return false
}
func (nl numList) FindSet(sum int) numList {
	for i, a := range nl {
		testSum := a
		for j := i + 1; j < len(nl); j++ {
			testSum += nl[j]
			if testSum > sum {
				break
			}
			if testSum == sum {
				return nl[i : j+1]
			}
		}
	}
	return nil
}
func (nl numList) Min() int {
	min := nl[0]
	for _, n := range nl[1:] {
		if n < min {
			min = n
		}
	}
	return min
}
func (nl numList) Max() int {
	max := nl[0]
	for _, n := range nl[1:] {
		if n > max {
			max = n
		}
	}
	return max
}

func loadData(fileName string) numList {
	debug("Loading data...")
	file, err := os.Open(fileName)
	if err != nil {
		exit(fmt.Sprintf("could not open %s!", fileName))
	}
	defer file.Close()
	s := bufio.NewScanner(file)
	var values numList

	for s.Scan() {
		line := s.Text()
		i := parseInt(line)
		values = append(values, i)
	}

	return values
}

func debug(s string) {
	fmt.Printf("DEBUG: %s\n", s)
}
func parseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		exit(fmt.Sprintf("Fatial: could not parse %s as int!", s))
	}

	return i
}
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
