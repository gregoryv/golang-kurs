# Exempel 6

Datatyper, siffror och sant eller falskt.

## main.go

	package main

	import "fmt"

	func main() {

	    var a string = "hello"

        fmt.Printf("%r\n", a)
        fmt.Printf("%r\n", 1)
        fmt.Printf("%r\n", 1.0)
        fmt.Printf("%r\n", true)
        fmt.Printf("%r\n", false)
	}

Kom ihåg att spara filen innan du fortsätter. Kör programmet genom att i terminalen skriva

	go run main.go

## Lärdom

1. `fmt.Printf` används när man vill formatera utskrift av olika värden
1. `%r` i första argumentet betyder, ersätt med datatyp och värde
1. `\n` betyder ny rad
1. typen för en variabel skrivs efter variabel deklarationen och före eventuell tilldelning


## Uppgift

1. Deklarera variabler för de övriga värdena som b, c, d och e med rätt datatyp och skriv ut dessa på samma sätt

[Exempel 7](../7/README.md#exempel-7)
