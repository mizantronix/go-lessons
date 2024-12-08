package main

import (
	"calculator/calc"
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage: cacl <command> <arg1> <arg2>")
		return
	}

	command := args[1]

	switch command {
	case "add":
		a, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Printf("Invalid number '%s'\n", args[2])
			return
		}

		b, err := strconv.Atoi(args[3])
		if err != nil {
			fmt.Printf("Invalid number '%s'\n", args[3])
			return
		}

		result := calc.Add(a, b)
		fmt.Printf("%d + %d = %d\n", a, b, result)

	case "substract":
		a, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Printf("Invalid number '%s'\n", args[2])
			return
		}

		b, err := strconv.Atoi(args[3])
		if err != nil {
			fmt.Printf("Invalid number '%s'\n", args[3])
			return
		}

		result := calc.Subtraction(a, b)
		fmt.Printf("%d - %d = %d\n", a, b, result)

	case "multiply":
		a, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Printf("Invalid number '%s'\n", args[2])
			return
		}

		b, err := strconv.Atoi(args[3])
		if err != nil {
			fmt.Printf("Invalid number '%s'\n", args[3])
			return
		}

		result := calc.Multiplication(a, b)
		fmt.Printf("%d * %d = %d\n", a, b, result)
	default:
		fmt.Println("Unknown command: '%s'", command)
	}
}
