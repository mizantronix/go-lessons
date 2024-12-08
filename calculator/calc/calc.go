package calc

func Factorial(n int) int {
	var result = 1
	for i := 1; i <= n; i++ {
		result *= i
	}

	return result
}

func Add(x, y int) int {
	return x + y
}

func Subtraction(x, y int) int {
	return x - y
}

func Multiplication(x, y int) int {
	return x * y
}

func Divide(x, y int) int {
	return x / y
}
