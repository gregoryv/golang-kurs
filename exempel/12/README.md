# Exempel 12

Interaktiva program

## main.go

    package main

	import (
        "bufio"
        "fmt"
        "os"
	)

	func main() {
        reader := bufio.NewReader(os.Stdin)
        fmt.Print("Enter text: ")
        text, _ := reader.ReadString('\n')
        fmt.Println(text)
	}


Kom ihåg att spara filen innan du fortsätter. Kör programmet genom att i terminalen skriva

	go run main.go

## Lärdom

1. Interaktiva program ger dig möjlighet att påverka ett program under tiden det körs
1. `os` paketet innehåller operativ system oberoende funktioner
1. I `os.Stdin` hamnar allt som användaren skriver under tiden ditt program körs
1. `bufio` är paketet som låter oss läsa vad användaren skriver

## Uppgift

1. Skriv om programmet så att det frågar efter två tal a och b och sen skriver ut summan av talen. Tips, du måste konvertera den inlästa strängen till ett heltal först. Till din hjälp har du [strconv.ParseInt](https://golang.org/pkg/strconv/#ParseInt) funktionen.
