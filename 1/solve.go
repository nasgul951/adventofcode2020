package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		exit("could not open input.txt!")
	}
	defer file.Close()
	var values []int
	s := bufio.NewScanner(file)

	i := 0
	for s.Scan() {
		i++
		line := s.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			exit(fmt.Sprintf("%s on line %d could not be converte to int.", line, i))
		}
		values = append(values, num)
	}

	for _, a := range values {
		for _, b := range values {
			for _, c := range values {
				if a+b+c == 2020 {
					fmt.Printf("%d + %d + %d = 2020\n", a, b, c)
					fmt.Printf("%d x %d x %d = %d\n", a, b, c, (a * b * c))
				}
			}
		}
	}

}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
