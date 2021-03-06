package main

import (
	"errors"
	"math"
	"regexp"
	"strconv"
)

type RegexpChecker struct {
	patternString string
}

func (r *RegexpChecker) Check(s string) bool {
	p, _ := regexp.Compile(r.patternString)
	return p.MatchString(s)
}

// Convert binary to decimal
func Bin2Dec(bin string) int {
	binaryFormatChecker := RegexpChecker{"^[0-1]+$"}
	if !binaryFormatChecker.Check(bin) {
		panic(errors.New("Bin2Dec, invalid input"))
	}
	// Using Library
	// v, _ := strconv.ParseInt(bin, 2, 0)
	// return int(v)

	// Manual
	dec := 0
	for i := 0; i < len(bin); i++ {
		v, _ := strconv.Atoi(bin[i : i+1])
		dec += v * int(math.Pow(2, float64(len(bin)-i-1)))
	}
	return dec
}
