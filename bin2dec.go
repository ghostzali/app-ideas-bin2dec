package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type RegexpChecker struct {
	patternString string
}

func (r *RegexpChecker) Check(s string) bool {
	p, _ := regexp.Compile(r.patternString)
	return p.MatchString(s)
}

// Convert binary to decimal
func Bin2Dec(bin string) (dec int) {
	binaryFormatChecker := RegexpChecker{"^[0-1]+$"}
	if !binaryFormatChecker.Check(bin) {
		panic(errors.New("Bin2Dec, invalid input"))
	}
	// Using Library
	v, _ := strconv.ParseInt(bin, 2, 0)
	dec = int(v)

	// Manual
	// for i := 0; i < len(bin); i++ {
	// 	v, _ := strconv.Atoi(bin[i : i+1])
	// 	dec += v * int(math.Pow(2, float64(len(bin)-i-1)))
	// }
	return
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	binaryFormatChecker := RegexpChecker{"^[0-1]+$"}
	for {
		fmt.Print(">> ")
		input, _ := reader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1)
		if input == "exit" {
			break
		}
		valid := binaryFormatChecker.Check(input)
		if !valid {
			fmt.Println("Invalid input!")
			continue
		}
		fmt.Printf("Binary %q equal to Decimal \"%d\"\n", input, Bin2Dec(input))
	}
}
