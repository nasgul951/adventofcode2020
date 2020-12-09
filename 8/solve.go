package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	insts := loadData("input.txt")

	fmt.Println("**************************************************")
	fmt.Println(" BEGIN TEST                                      *")
	fmt.Println("**************************************************")
	fmt.Println("")

	passCounter := 0
	for {
		acc, p, op, inTest := 0, 0, "", false
		for {
			i := insts[p]
			if i.Count > 0 {
				fmt.Printf(" *** Repeated Instruction\n")
				break
			}

			fmt.Printf("%03d: %s", p, i.Op)

			op, inTest = getTestOp(i, inTest)

			fmt.Printf(" %d", i.Num)

			switch op {
			case "acc":
				acc += i.Num
				p++
			case "jmp":
				p += i.Num
			case "nop":
				p++
			default:
				exit(fmt.Sprintf("Invalid instruction '%s'!", i.Op))
			}
			i.Count++
			fmt.Printf("\n")

			if p == len(insts) {
				exit(fmt.Sprintf("Reached end of program in %d passes, Acc: %d", passCounter, acc))
			}

			if p > len(insts) {
				exit("Failure, p has exceded len of instructions.")
			}
		}
		insts.ResetCounters()
		passCounter++

		if !inTest {
			exit("Completed a pass without testing anything!!")
		}
	}
}

func getTestOp(i *instruction, inTest bool) (string, bool) {
	if inTest {
		fmt.Printf(" -- In Test -- ")
		return i.Op, true
	}

	if !i.Tested && i.Op != "acc" {
		i.Tested = true
		if i.Op == "nop" {
			fmt.Printf("->jmp")
			return "jmp", true
		} else if i.Op == "jmp" {
			fmt.Printf("->nop")
			return "nop", true
		}
	} else if i.Tested && i.Op != "aac" {
		fmt.Printf("*")
	}

	return i.Op, false
}

type instruction struct {
	Op     string
	Sign   string
	Num    int
	Count  int
	Tested bool
}
type instList []*instruction

func newInstruction(s string) *instruction {
	r := regexp.MustCompile(`(?P<op>[a-z]{3})\s(?P<num>[\+|-][0-9]+)`)
	m := r.FindStringSubmatch(s)

	i := new(instruction)
	i.Op = m[1]
	i.Num = parseInt(m[2])
	i.Count = 0
	i.Tested = false

	return i
}

func (il instList) ResetCounters() int {
	c := 0
	for _, i := range il {
		i.Count = 0
		if i.Tested {
			c++
		}
	}
	return c
}

func loadData(fileName string) instList {
	debug("Loading data...")
	file, err := os.Open(fileName)
	if err != nil {
		exit(fmt.Sprintf("could not open %s!", fileName))
	}
	defer file.Close()
	s := bufio.NewScanner(file)
	var values instList

	for s.Scan() {
		line := s.Text()
		i := newInstruction(line)
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
