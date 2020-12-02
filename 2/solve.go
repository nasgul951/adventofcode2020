package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		exit("could not open input.txt!")
	}
	defer file.Close()
	s := bufio.NewScanner(file)

	i := 0
	goodCount := 0
	badCount := 0
	for s.Scan() {
		i++
		line := s.Text()
		r := parseRecord(line)

		if isGoodv2(r) {
			goodCount++
		} else {
			fmt.Printf("%s is not valid\n", line)
			badCount++
		}
	}

	fmt.Printf("%d bad, %d good, Total Records Processed: %d\n", badCount, goodCount, i)
}

// Record - propably don't need to export this.
type Record struct {
	Password string
	Required byte
	Min      int
	Max      int
}

func isGood(record Record) bool {
	count := countLetter(record.Password, record.Required)
	if (count >= record.Min) && (count <= record.Max) {
		return true
	}

	return false
}

func isGoodv2(record Record) bool {
	var bytes = []byte(record.Password)
	match1 := bytes[record.Min-1] == record.Required
	match2 := bytes[record.Max-1] == record.Required
	if (match1 || match2) && !(match1 && match2) {
		return true
	}

	return false
}

func parseRecord(source string) Record {
	var record Record

	val := strings.Split(source, ":")
	policy := strings.Split(val[0], " ")
	valueRange := strings.Split(policy[0], "-")
	record.Password = strings.TrimSpace(val[1])
	record.Required = policy[1][0]
	record.Min, _ = strconv.Atoi(valueRange[0])
	record.Max, _ = strconv.Atoi(valueRange[1])

	return record
}

func countLetter(source string, letter byte) int {
	counter := 0
	bytes := []byte(source)
	for _, b := range bytes {
		if b == letter {
			counter++
		}
	}

	return counter
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
