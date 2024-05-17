package cli

import (
	"fmt"
	"strconv"
	"strings"
)

var CHOICES = []string{"Rate Names", "Show results", "Compare Results"}

func CheckForType() int {
	var currentInput string
	asterisks := strings.Repeat("*", 50)
	for {
		fmt.Printf("\nWhat do you want to do?\n%s\n", asterisks)
		for i, choice := range CHOICES {
			fmt.Printf("%d: %s\n", i+1, choice)
		}
		fmt.Printf("Enter the number: ")
		fmt.Scanln(&currentInput)
		rank, err := strconv.Atoi(currentInput)
		if err == nil && rank >= 1 && rank <= len(CHOICES) {
			return rank
		}
		fmt.Printf("Invalid input. Please enter a number between 1 and %d.", len(CHOICES))
	}
}
