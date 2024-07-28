package main

import (
	"fmt"
	"strconv"
)

// Function to validate the credit card number using Luhn algorithm
func validateCard(number string) bool {
	// Remove all whitespaces from the number
	number = removeSpaces(number)

	// Check if the card number length is valid based on card type (basic check)
	if !isValidLength(number) {
		return false
	}

	// Initialize variables for Luhn algorithm
	sum := 0
	isDouble := false

	// Loop through each digit of the number (from right to left)
	for i := len(number) - 1; i >= 0; i-- {
		digit, err := strconv.Atoi(string(number[i])) // Convert rune to int
		if err != nil {
			return false // Invalid character in number
		}

		// Multiply digit based on its position (double every other digit)
		var newDigit int
		if isDouble {
			newDigit = digit * 2
			if newDigit > 9 {
				newDigit = newDigit - 9
			}
		} else {
			newDigit = digit
		}

		// Update sum and double flag
		sum += newDigit
		isDouble = !isDouble
	}

	// Check if the sum is a multiple of 10 (Luhn algorithm validation)
	return sum%10 == 0
}

// Function to remove spaces from the card number
func removeSpaces(str string) string {
	var newString string
	for _, char := range str {
		if char != ' ' {
			newString += string(char)
		}
	}
	return newString
}

// Function to perform basic check on card number length based on type
func isValidLength(number string) bool {
	// Check for common card lengths (Visa, Mastercard, Amex, Discover)
	lengths := map[string]int{"Visa": 16, "Mastercard": 16, "Amex": 15, "Discover": 16}
	for _, length := range lengths {
		if len(number) == length {
			return true
		}
	}
	return false
}

func main() {
	// Test the function
	fmt.Println(validateCard("4532 3195 8579 2454")) // Should return true
	fmt.Println(validateCard("6011 1111 1111 1117")) // Should return true
	fmt.Println(validateCard("3782 8224 6310 005"))  // Should return true
	fmt.Println(validateCard("6011 1111 1111 1118")) // Should return false
}
