package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	seats := loadData()
	fmt.Printf("Max Seat ID: %d", seats.MaxID())
}

type seat struct {
	Row    int
	Column int
	ID     int
}
type seatList []seat

func (l seatList) MaxID() int {
	max := 0
	for _, v := range l {
		if v.ID > max {
			max = v.ID
		}
	}
	return max
}

func getSeat(location string) seat {
	var s seat
	s.Row = partition(location[0:7], 0, 127)
	s.Column = partition(location[7:10], 0, 7)
	s.ID = s.Row*8 + s.Column

	return s
}

func partition(s string, min, max int) int {
	fmt.Printf("DEBUG: partitioning '%s'\n", s)
	for _, c := range s {
		mid := (max - min) / 2
		if string(c) == "F" || string(c) == "L" {
			max = min + mid
		} else if string(c) == "B" || string(c) == "R" {
			min = min + mid + 1
		} else {
			exit(fmt.Sprintf("Unexpected char %c", c))
		}
		fmt.Printf("DEBUG: %q: %d - %d\n", c, min, max)
		if min == max {
			return min
		}
	}
	exit(fmt.Sprintf("ERROR: failed to partition '%s' final result: min(%d) - max(%d)\n", s, min, max))
	return -1
}

func loadData() seatList {
	file, err := os.Open("input.txt")
	if err != nil {
		exit("could not open input.txt!")
	}
	defer file.Close()
	s := bufio.NewScanner(file)
	var values seatList

	for s.Scan() {
		line := s.Text()
		values = append(values, getSeat(line))
	}

	return values
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
