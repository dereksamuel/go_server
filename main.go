package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin) // input's terminal getting
	fmt.Println("I/")
	scanner.Scan()
	operation := scanner.Text() // asignar al scaner, y se obtiene como text ya que es un buffer
	fmt.Println(operation)
	fmt.Println("/O")

	values := strings.Split(operation, "+") // parte al texto
	fmt.Println(values)
}
