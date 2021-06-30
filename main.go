package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	binaryFormatChecker := RegexpChecker{"^[0-1]{1,8}$"}
	for {
		fmt.Println("Choose operation mode:")
		fmt.Println("1: Binary to Decimal")
		fmt.Println("2: Decimal to Binary")
		fmt.Println("exit: Close the program")
		fmt.Print(">> ")
		option, _ := reader.ReadString('\n')
		option = strings.ReplaceAll(option, "\n", "")
		if option == "exit" {
			break
		}
		switch option {
		case "1":
			fmt.Println("Input binary digits (0/1) up to 8 digits:")
			fmt.Print(">> ")
			input, _ := reader.ReadString('\n')
			input = strings.ReplaceAll(input, "\n", "")
			valid := binaryFormatChecker.Check(input)
			if !valid {
				fmt.Println("Invalid input! Please enter binary digits (0/1) up to 8 digits")
				continue
			}
			fmt.Printf("Binary %q equal to Decimal \"%d\"\n", input, Bin2Dec(input))
		case "2":
			fmt.Println("Input decimal number")
			fmt.Print(">> ")
			input, _ := reader.ReadString('\n')
			input = strings.ReplaceAll(input, "\n", "")
			dec, err := strconv.ParseInt(input, 10, 0)
			if err != nil {
				fmt.Println("Invalid input! Please enter decimal number")
				continue
			}
			fmt.Printf("Decimal \"%d\" equal to Binary %q\n", dec, Dec2Bin(int(dec)))
		default:
			continue
		}
	}
}
