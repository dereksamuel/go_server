package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin) // input's terminal getting
	fmt.Println("I/")
	scanner.Scan()
	operation := scanner.Text() // asignar al scaner, y se obtiene como text ya que es un buffer
	fmt.Println(operation)
	fmt.Println("/O")

	operator := "-"
	values := strings.Split(operation, operator) // parte al texto
	fmt.Println(values)

	result := 0

	if operator == "*" || operator == "/" {
		result = 1
	}

	for i, val := range values {
		value, errorResult := strconv.Atoi(val)
		thereIsAnError := false

		if errorResult != nil {
			fmt.Println("[error]:", errorResult)
			break
		}

		switch operator {
		case "+":
			result = result + value
		case "*":
			result = result * value
		default:
			fmt.Println("[error]: No esta soportado el operador:", operator)
			thereIsAnError = true
		}

		if thereIsAnError != false {
			break
		}

		if i+1 == len(values) {
			fmt.Println("[result OK]:", result)
		}
	}
}
