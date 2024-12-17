package calculation

import (
	"strconv"
	"unicode"
)

func PriorityOperation(operation string) int {
	switch operation {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	}
	return 0
}

func ArithmeticOperation(a, b float64, operation string) (float64, error) {
	switch operation {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, ErrDivisionByZero
		}
		return a / b, nil
	}
	return 0, nil
}

func Calc(expression string) (float64, error) {

	var numbers []float64
	var operations []string
	var numberBuffer string

	for _, str := range expression {
		if unicode.IsDigit(str) || str == '.' {
			numberBuffer += string(str)
		} else {
			if numberBuffer != "" {
				n, err := strconv.ParseFloat(numberBuffer, 64)
				if err != nil {
					return 0, ErrInvalidExpression
				}
				numbers = append(numbers, n)
				numberBuffer = ""
			}
			if str == '+' || str == '-' || str == '*' || str == '/' {
				for len(operations) > 0 && PriorityOperation(operations[len(operations)-1]) >= PriorityOperation(string(str)) {
					if len(numbers) < 2 {
						return 0, ErrInvalidExpression
					}
					n2 := numbers[len(numbers)-1]
					numbers = numbers[:len(numbers)-1]
					n1 := numbers[len(numbers)-1]
					numbers = numbers[:len(numbers)-1]
					operation := operations[len(operations)-1]
					operations = operations[:len(operations)-1]
					result, err := ArithmeticOperation(n1, n2, operation)
					if err != nil {
						return 0, err
					}
					numbers = append(numbers, result)
				}
				operations = append(operations, string(str))
			} else if str == '(' {
				operations = append(operations, string(str))
			} else if str == ')' {
				for len(operations) > 0 && operations[len(operations)-1] != "(" {
					if len(numbers) < 2 {
						return 0, ErrInvalidExpression
					}
					n2 := numbers[len(numbers)-1]
					numbers = numbers[:len(numbers)-1]
					n1 := numbers[len(numbers)-1]
					numbers = numbers[:len(numbers)-1]
					operation := operations[len(operations)-1]
					operations = operations[:len(operations)-1]
					result, err := ArithmeticOperation(n1, n2, operation)
					if err != nil {
						return 0, err
					}
					numbers = append(numbers, result)
				}
				if len(operations) == 0 {
					return 0, ErrInvalidExpression
				}
				operations = operations[:len(operations)-1]
			}
		}
	}

	if numberBuffer != "" {
		n, err := strconv.ParseFloat(numberBuffer, 64)
		if err != nil {
			return 0, ErrInvalidExpression
		}
		numbers = append(numbers, n)
	}

	for len(operations) > 0 {
		if len(numbers) < 2 {
			return 0, ErrInvalidExpression
		}
		n2 := numbers[len(numbers)-1]
		numbers = numbers[:len(numbers)-1]
		n1 := numbers[len(numbers)-1]
		numbers = numbers[:len(numbers)-1]
		operation := operations[len(operations)-1]
		operations = operations[:len(operations)-1]
		result, err := ArithmeticOperation(n1, n2, operation)
		if err != nil {
			return 0, err
		}
		numbers = append(numbers, result)
	}

	if len(numbers) != 1 {
		return 0, ErrInvalidExpression
	}
	return numbers[0], nil
}
