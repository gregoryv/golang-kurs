package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func askNumber() int64 {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	line, _ := reader.ReadString('\n')
	noEnter := strings.Replace(line, "\n", "", 1)
	number, _ := strconv.ParseInt(noEnter, 10, 64)
	return number
}

func main() {
	a := askNumber()
	b := askNumber()
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
}
