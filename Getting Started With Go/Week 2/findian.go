package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
)

func main() {
	var val1, val2, val3 int
	fmt.Printf("\nEnter String To Check: ")
	reader := bufio.NewReader(os.Stdin) // Had to use Stdin as input after spaces were ignored
	usr_inp, _ := reader.ReadString('\n')
	final_str :=strings.Trim(usr_inp, "\n") // Removing '\n' as it was giving a lot of random errors
	final_str = strings.ToLower(final_str)
	first_char := final_str[0:1]
	last_char := final_str[len(final_str)-1:] 
	// converting into boolean logic
	if first_char == "i" {
		val1 = 1
	} else {
		val1 = 0
	}
	if last_char == "n" {
		val2 = 1
	} else {
		val2 = 0
	}
	
	chk_str := strings.ContainsAny(final_str, "a") // checking string if it contains "a"
	
	if chk_str {
		val3 = 1
	} else {
		val3 = 0
	}
// Boolean Check if First Letter Contains "i", ends with "n" and contains "a"
	if val1 == 1 && val2 == 1 && val3 == 1 {
		fmt.Printf("\nFound!\n")
	} else {
		fmt.Print("\nNot Found!\n")
	}
}