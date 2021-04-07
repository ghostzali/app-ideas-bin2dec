package main

import (
	"math"
	"regexp"
)

func Dec2Bin(dec int) string {
	s := ""
	remain := dec
	for i := 31; i >= 0; i-- {
		v := int(math.Pow(2, float64(i)))
		if remain-v >= 0 {
			s += "1"
			remain -= v
		} else {
			s += "0"
		}
	}
	r, _ := regexp.Compile("^0+")
	return r.ReplaceAllString(s, "")

	// Using Library
	// return strconv.FormatInt(int64(dec), 2)
}
