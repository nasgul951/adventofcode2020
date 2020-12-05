package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	boardingPasses := loadData()
	allSeatIds := getAllSeatIds()
	sort.Ints(allSeatIds)

	for _, s := range allSeatIds {
		_, found := boardingPasses.GetByID(s)
		if found {
			fmt.Printf("%d ***\n", s)
		} else {
			fmt.Printf("%d\n", s)
		}
	}
	fmt.Printf("Max Seat ID: %d", boardingPasses.MaxID())
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
func (l seatList) GetByID(id int) (seat, bool) {
	for _, s := range l {
		if s.ID == id {
			return s, true
		}
	}
	var e seat
	return e, false
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

func getAllSeatIds() []int {
	var seatIds []int
	for r := 0; r < 128; r++ {
		for c := 0; c < 8; c++ {
			seatIds = append(seatIds, (r*8 + c))
		}
	}

	return seatIds
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
