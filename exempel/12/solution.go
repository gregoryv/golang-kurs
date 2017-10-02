package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func askNumber() int64 {
	// Ask a question
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter number: ")
	// Read the answer
	line, _ := reader.ReadString('\n')
	// Remove \n characters at end of line
	noEnter := strings.Replace(line, "\n", "", 1)
	// Convert the string to a number
	number, _ := strconv.ParseInt(noEnter, 10, 64)
	return number
}

func main() {
	a := askNumber()
	b := askNumber()
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
}
