package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	calculator1()
}

func calculator1() {
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	multiWriter := io.MultiWriter(os.Stdout, file)
	fmt.Fprintln(file, "Calculator started")

	for {
		arg1 := promptNum("Please input first argument")
		operator := promptStr("Please input operator")
		arg2 := promptNum("Please input second argument")

		result, err := calc(arg1, operator, arg2)
		if err != nil {
			fmt.Fprintln(multiWriter, err)
		} else {
			fmt.Fprintf(multiWriter, "%d %s %d = %f\n", arg1, operator, arg2, result)
		}

		continueWill := promptStr("Do you want to continue? (y/n)")
		if strings.ToLower(continueWill) == "n" {
			break
		}
	}

}

func promptNum(prompt string) int {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println(prompt)
		scanner.Scan()
		input := scanner.Text()
		n, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid number, please enter a valid integer.")
			continue
		}
		return n
	}
}

func promptStr(prompt string) string {
	var s string
	fmt.Println(prompt)
	fmt.Scanln(&s)
	return s
}

func calc(arg1 int, operator string, arg2 int) (float64, error) {
	switch operator {
	case "+":
		return float64(arg1 + arg2), nil
	case "-":
		return float64(arg1 - arg2), nil
	case "*":
		return float64(arg1 * arg2), nil
	case "/":
		if arg2 == 0 {
			return 0, fmt.Errorf("cannot divide by zero")
		}
		return float64(arg1) / float64(arg2), nil
	default:
		return 0, fmt.Errorf("invalid operator: %s", operator)
	}
}
