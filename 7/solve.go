package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	bags := loadData()

	bagColor := "shiny gold"
	bagCount := 0
	for _, b := range bags {
		printBag(b)
		holdCount := b.CanHold(bagColor, 0)
		fmt.Printf("Can Hold: %d %s.\n", holdCount, bagColor)
		if holdCount > 0 {
			bagCount++
		}
	}

	fmt.Printf("%d bags can hold %s.\n", bagCount, bagColor)

	myBag, found := bags.GetByColor(bagColor)
	if !found {
		exit("could not find my bag.")
	}
	fmt.Printf("%s contains %d other bags.\n", bagColor, myBag.CountInnerBags(1, 0))
}

func printBag(b *bag) {
	fmt.Printf("%s contains", b.Color)
	for _, c := range b.Bags {
		fmt.Printf(", %d %s", c.Qty, c.Bag.Color)
	}
	fmt.Printf("\n")
}

type bag struct {
	Source   string
	Color    string
	Contains string
	Bags     []bagContainer
}
type bagContainer struct {
	Qty int
	Bag *bag
}
type bagList []*bag

func (b *bag) CanHold(color string, depth int) int {
	if depth > 30 {
		exit("Too much recursion, exiting.")
	}
	counter := 0
	for _, c := range b.Bags {
		if c.Bag.Color == color || color == "" {
			counter += (c.Qty * depth)
		}
		counter += c.Bag.CanHold(color, depth+1)
	}

	return counter
}
func (b *bag) CountInnerBags(q int, depth int) int {
	if depth > 30 {
		exit("Too much recursion, exiting.")
	}
	counter := 0
	for _, c := range b.Bags {
		counter += c.Qty * q
		counter += c.Bag.CountInnerBags(c.Qty*q, depth+1)
	}

	return counter
}
func (bl bagList) GetByColor(c string) (*bag, bool) {
	for _, b := range bl {
		if b.Color == c {
			return b, true
		}
	}
	return nil, false
}
func (bl bagList) UpdateLinks() {
	for _, b := range bl {
		debug(fmt.Sprintf("UpdteLinks for: %s", b.Color))
		list := strings.Split(b.Contains, ",")

		r := regexp.MustCompile(`(?P<qty>[0-9]+)\s(?P<color>.+)\sbag`)
		for _, c := range list {
			m := r.FindStringSubmatch(c)
			if len(m) != 3 {
				if c == "no other bags." {
					continue
				} else {
					exit(fmt.Sprintf("could not parse %s, only found %d matches", c, len(m)))
				}
			}

			qty := parseInt(m[1])
			color := m[2]

			b1, found := bl.GetByColor(color)
			if !found {
				exit(fmt.Sprintf("Did not find bag color:%s", color))
			}

			bagLink := new(bagContainer)
			bagLink.Qty = qty
			bagLink.Bag = b1
			b.Bags = append(b.Bags, *bagLink)
			debug(fmt.Sprintf("bagLink: %d %s", bagLink.Qty, bagLink.Bag.Color))
		}

	}
}
func newBag(s string) *bag {
	r := regexp.MustCompile(`(?P<color>.+)\sbags\scontain\s(?P<contains>.*)`)
	m := r.FindStringSubmatch(s)

	b := new(bag)
	b.Source = s
	b.Color = m[1]
	b.Contains = m[2]

	return b
}

func loadData() bagList {
	debug("Loading data...")
	file, err := os.Open("input.txt")
	if err != nil {
		exit("could not open input.txt!")
	}
	defer file.Close()
	s := bufio.NewScanner(file)
	var values bagList

	for s.Scan() {
		line := s.Text()
		debug(fmt.Sprintf("read Line: %s", line))
		b := newBag(line)
		values = append(values, b)
	}

	values.UpdateLinks()

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
