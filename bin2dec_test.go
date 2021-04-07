package main

import (
	"fmt"
	"testing"
)

func TestBin2Dec(t *testing.T) {
	got := Bin2Dec("1011")
	want := 11

	if got != want {
		t.Errorf("Bin2Dec result got %d, want %d", got, want)
	}
}

func ExampleBin2Dec() {
	fmt.Printf("Binary '1101' equal to decimal '%d'", Bin2Dec("1101"))

	// Output: Binary '1101' equal to decimal '13'
}

func BenchmarkBin2Dec(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bin2Dec("10100101")
	}
}
