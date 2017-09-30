# Exempel 2

Använda andra paket.

## main.go

	package main

	import "fmt"

	func main() {
		fmt.Println("Hello, World!")
	}


Kom ihåg att spara filen innan du fortsätter.
Kör programmet genom att i terminalen skriva

	go run main.go

## Lärdom

1. `import` används för att återanvända kod i andra `paket`
1. `fmt.Println` funktionsanrop till andra paket
1. publika funktioner i paket börjar med Versaler, privata med gemener
1. vilka paket finns? det finns massor, https://golang.org/pkg beskriver de som kommer med språket

[Exempel 3](../3/README.md#exempel-3)
