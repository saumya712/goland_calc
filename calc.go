package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func add(a int64, b int64) int64 {
	return a + b
}

func sub(a int64, b int64) int64 {
	return a - b
}

func mul(a int64, b int64) int64 {
	return a * b
}

func div(a int64, b int64) int64 {
	return a / b
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter number 1:")
	input1, _ := reader.ReadString('\n')
	input1 = strings.TrimSpace(input1)
	num1, _ := strconv.ParseInt(input1, 10, 64)

	fmt.Println("Enter number 2:")
	input2, _ := reader.ReadString('\n')
	input2 = strings.TrimSpace(input2)
	num2, _ := strconv.ParseInt(input2, 10, 64)

	fmt.Println("Select the operation you want to perform (+, -, *, /):")
	opp, _ := reader.ReadString('\n')
	opp = strings.TrimSpace(opp)

	if opp == "+" {
		result := add(num1, num2)
		fmt.Println("Result:", result)
	} else if opp == "-" {
		result := sub(num1, num2)
		fmt.Println("Result:", result)
	} else if opp == "*" {
		result := mul(num1, num2)
		fmt.Println("Result:", result)
	} else if opp == "/" {
		if num2 == 0 {
			fmt.Println("Error: Cannot divide by zero")
		} else {
			result := div(num1, num2)
			fmt.Println("Result:", result)
		}
	} else {
		fmt.Println("Invalid operation")
	}
}
