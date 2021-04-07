package main

import (
	"fmt"
	"testing"
)

func TestDec2Bin(t *testing.T) {
	got := Dec2Bin(11)
	want := "1011"

	if got != want {
		t.Errorf("Dec2Bin result got %q, want %q", got, want)
	}
}

func ExampleDec2Bin() {
	fmt.Printf("Decimal '13' equal to binary '%s'", Dec2Bin(13))

	// Output: Decimal '13' equal to binary '1101'
}

func BenchmarkDec2Bin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Dec2Bin(2512)
	}
}
