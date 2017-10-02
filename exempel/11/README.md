# Exempel 11

Listor och loopar

## main.go

    package main

	import "fmt"

	func main() {
	    x := []int{0,0,0,0,0,0}
	    fmt.Println(x)
	    x[1] = 9
	    fmt.Println(x)
	    for index, value := range x {
	        fmt.Printf("%d. %d\n", index, value)
	    }
	}

Kom ihåg att spara filen innan du fortsätter. Kör programmet genom att i terminalen skriva

	go run main.go

## Lärdom

1. Listor av värden deklareras med hakparenteser före datatypen `[]`
1. Listor indexeras med första index på 0
1. Nyckelordet `for` används för att gå igenom(loop över) en sekvens av värden
1. `range` används för att stega igenom en indexerad lista

## Uppgift

1. Skriv en funktion `sumAll(numbers []int) int` som räknar ut summan av alla heltal i en lista
1. Skriv en funktion `countEven(numbers []int) int` som räknar antalet heltal i en lista

I andra uppgiften skall main se ut

    func main() {
	    numbers := []int{1,2,3,4,5}
	    even = countEven(numbers)
	    fmt.Println(numbers)
	    fmt.Printf("has %d even numbers\n", even)
	}

*[Lösningsförslag för båda uppgifterna](./solution.go)*

[Exempel 12](../12/README.md#exempel-12)
