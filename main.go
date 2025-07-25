package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
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
			logf(multiWriter, "Error: %v", err)
		} else {
			fmt.Printf("%d %s %d = %f\n", arg1, operator, arg2, result)
			logf(file, "%d %s %d = %f", arg1, operator, arg2, result)
		}

		continueWill := promptStr("Do you want to continue? (y/n)")
		if strings.ToLower(continueWill) == "n" {
			logf(multiWriter, "Calculator exited")
			break
		}
	}
}

func logf(w io.Writer, format string, a ...interface{}) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Fprintf(w, "[%s] %s\n", timestamp, fmt.Sprintf(format, a...))
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
