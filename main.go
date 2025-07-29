package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

func main() {
	// fmt.Println("浮動小数点数の計算の不正確性")
	// a := 0.1
	// b := 0.2
	// fmt.Printf("a = %.17f\n", a)
	// fmt.Printf("b = %.17f\n", b)
	// fmt.Printf("a + b = %.17f\n", a+b)
	// fmt.Printf("Expected: 0.30000000000000000\n\n")

	calculator()
}

func calculator() {
	// Open log file in append mode
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	multiWriter := io.MultiWriter(os.Stdout, file)
	logf(file, "Calculator started")

	for {
		arg1 := promptNum("Please input first argument")
		operator := promptStr("Please input operator")
		arg2 := promptNum("Please input second argument")

		result, err := calc(arg1, operator, arg2)
		if err != nil {
			logf(multiWriter, fmt.Sprintf("Error: %v", err))
		} else {
			fmt.Printf("%v %s %v = %v\n", arg1, operator, arg2, result)
			logf(file, fmt.Sprintf("%v %s %v = %v", arg1, operator, arg2, result))
		}

		continueWill := promptStr("Do you want to continue? (y/n)")
		if strings.ToLower(continueWill) == "n" {
			logf(multiWriter, "Calculator exited")
			break
		}
	}
}

func logf(w io.Writer, content string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Fprintf(w, "[%s] %s\n", timestamp, content)
}

func promptNum(prompt string) decimal.Decimal {
	input := promptStr(prompt)
	d, err := decimal.NewFromString(input)
	if err != nil {
		fmt.Println("Invalid number, please enter a valid number.")
		return promptNum(prompt)
	}
	return d
}

func promptStr(prompt string) string {
	var s string
	fmt.Println(prompt)
	fmt.Scanln(&s)
	return s
}

func calc(arg1 decimal.Decimal, operator string, arg2 decimal.Decimal) (decimal.Decimal, error) {
	switch operator {
	case "+":
		return arg1.Add(arg2), nil
	case "-":
		return arg1.Sub(arg2), nil
	case "*":
		return arg1.Mul(arg2), nil
	case "/":
		if decimal.Zero.Equal(arg2) {
			return decimal.Zero, fmt.Errorf("cannot divide by zero")
		}
		return arg1.Div(arg2), nil
	default:
		return decimal.Zero, fmt.Errorf("invalid operator: %s", operator)
	}
}
