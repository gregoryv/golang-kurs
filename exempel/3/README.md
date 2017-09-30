# Exempel 3

Funktioner och argument.

## main.go

	package main

	import "fmt"

	func main() {
		display("Hello, World!")
	}

	func display(something string) {
		fmt.Println(something)
	}

Kom ihåg att spara filen innan du fortsätter. Kör programmet genom att i terminalen skriva

	go run main.go

## Lärdom

1. `display` är en egen funktion som tar ett argument `something`
1. `string` datatypen är en sträng, kompilatorn hjälper dig att skilja på olika datatyper.
1. `datatyper` kan vara sträng, tecken, siffra, decimaltal osv.

[Exempel 4](../4/README.md#exempel-4)
