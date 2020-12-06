package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	groups := loadData()

	// i := 44
	// for _, p := range groups[i].People {
	// 	fmt.Printf("%s\n", p)
	// }
	// fmt.Printf("\n%s (%d)\n", string(groups[i].YesAnswers), len(groups[i].YesAnswers))
	yesCounter := 0
	for _, g := range groups {
		for _, p := range g.People {
			fmt.Printf("%s\n", p)
		}
		fmt.Printf("everyone answered yes to %d in this group.\n", len(g.YesAnswers))
		yesCounter += len(g.YesAnswers)
	}
	fmt.Printf("%d Total Yes Answers.\n", yesCounter)
}

type group struct {
	People     []string
	YesAnswers answers
}
type groupList []*group
type answers []rune

func (g *group) AddPerson(s string) {
	if len(g.People) == 0 {
		g.YesAnswers.Add(s)
	} else {
		g.YesAnswers.Intersect(s)
	}
	g.People = append(g.People, s)
}

func (rl answers) Exists(r rune) bool {
	for _, r1 := range rl {
		if r1 == r {
			return true
		}
	}
	return false
}
func (rl *answers) Add(s string) {
	debug(fmt.Sprintf("adding %s to answers\n", s))
	for _, r := range s {
		*rl = append(*rl, r)
	}
}
func (rl *answers) Intersect(s string) {
	var a answers
	for _, r := range *rl {
		if strings.Contains(s, string(r)) {
			a = append(a, r)
		}
	}
	*rl = a
}

func loadData() groupList {
	file, err := os.Open("input.txt")
	if err != nil {
		exit("could not open input.txt!")
	}
	defer file.Close()
	s := bufio.NewScanner(file)
	var values groupList

	g := new(group)
	for s.Scan() {
		line := s.Text()
		if strings.TrimSpace(line) == "" {
			values = append(values, g)
			g = new(group)
		} else {
			g.AddPerson(line)
		}
	}
	values = append(values, g)

	return values
}

func debug(s string) {
	//fmt.Printf("DEBUG: %s\n", s)
}
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
