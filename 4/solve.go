package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		exit("could not open input.txt!")
	}
	defer file.Close()
	s := bufio.NewScanner(file)
	var values []map[string]string

	i := 0
	data := ""
	for s.Scan() {
		i++
		line := s.Text()
		if line != "" {
			data += line + " "
		} else {
			values = append(values, toMap(data))
			data = ""
		}
	}
	values = append(values, toMap(data))

	valid2Count := 0
	var test []string
	for _, p := range values {
		if isValidv2(p) {
			valid2Count++

			t := p["pid"]
			if strings.HasSuffix(t, "") {
				test = append(test, t)
			}
		}
	}
	fmt.Printf("test len:%d\n", len(test))
	sort.Strings(test)

	lastv := ""
	for _, v := range test {
		fmt.Printf("pid:'%s'\n", v)
		if lastv == v {
			fmt.Printf("*** DUPLICATE PID ***\n")
		}
		lastv = v
	}
	//224
	fmt.Printf("Found %d valid records.\n", valid2Count)
}

func toMap(data string) map[string]string {
	newMap := make(map[string]string)

	trimmed := strings.TrimSpace(data)
	list := strings.Split(trimmed, " ")
	for _, l := range list {
		k := strings.Split(l, ":")
		newMap[k[0]] = k[1]
	}

	return newMap
}

func isValidv1(p map[string]string) bool {
	required := [][]string{
		{"byr", "^19[2-9][0-9]|200[0-2]$"},
		{"iyr", "^201[0-9]|2020$"},
		{"eyr", "^202[0-9]|2030$"},
		{"hgt", "^1[5-8][0-9]cm|19[0-3]cm|59in|6[0-9]in|7[0-6]in$"},
		{"hcl", "^#[0-9a-f]$"},
		{"ecl", "^amb|blu|brn|gry|grn|hzl|oth$"},
		{"pid", "^[0-9]{9}$"},
		{"cid", ".*"},
	}

	for _, k := range required {
		_, exists := p[k[0]]
		if !exists {
			return false
		}
	}

	return true
}

func isValidv2(p map[string]string) bool {
	required := [][]string{
		{"byr", "^19[2-9][0-9]|200[0-2]$"},
		{"iyr", "^201[0-9]|2020$"},
		{"eyr", "^202[0-9]|2030$"},
		{"hgt", "^1[5-8][0-9]cm|19[0-3]cm|59in|6[0-9]in|7[0-6]in$"},
		{"hcl", "^#[0-9a-f]{6}$"},
		{"ecl", "^amb|blu|brn|gry|grn|hzl|oth$"},
		{"pid", "^[0-9]{9}$"},
	}

	for _, k := range required {
		v, exists := p[k[0]]
		if !exists {
			return false
		}
		re := regexp.MustCompile(k[1])
		if !re.MatchString(v) {
			return false
		}

	}

	return true
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
