# Exempel 10

Uttrycket `switch`

## main.go

    package main

	import "fmt"

	func main() {
	    i := 0
	    switch i {
            case 1: fmt.Print("one")
            case 2: fmt.Print("two")
            case 3, 4, 5: fmt.Print("three, four or five")
            default: fmt.Print(i, " is unknown")
	    }
	}

Kom ihåg att spara filen innan du fortsätter. Kör programmet genom att i terminalen skriva

	go run main.go

## Lärdom

1. `switch` uttrycket använder du för att göra ett av många val

## Uppgift

1. Flytta ut ovan switch uttryck till funktionen `numberIs(i int)`
2. Skriv en ny funktion `numberIsAlso(i int)` som istället använder `if` och `else` uttrycken för att åstadkomma samma resultat som `numberIs(i int)` funktionen. Anropa båda funktionerna och jämför svaren, t.ex

	    func main() {
	        i := 0
	        numberIs(i)
	        numberIsAlso(i)
	    }

3. Skriv om funktionerna `numberIs` och `numberIsAlso` så att dessa returnerar en sträng. Skriv main så den jämför dina svar, t.ex.

	    func main() {
	        i := 0
	        fmt.Println(numberIs(i), " = ", numberIsAlso(i))
	    }

[Exempel 11](../11/README.md#exempel-11)
