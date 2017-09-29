# Exempel 9

Funktioner med returvärden.

## main.go

    package main

	import "fmt"

	func sum(a, b int) int {
	    return a + b
	}

	func order(a, b int) (int, int) {
		if(a < b) {
			return a, b
		}
		return b,a
	}

	func main() {
	    answer := sum(1,2)
	    fmt.Println(answer)
		fmt.Println(order(9,4))
		fmt.Println(order(1,11))
	}

Kom ihåg att spara filen innan du fortsätter. Kör programmet genom att i terminalen skriva

	go run main.go

## Lärdom

1. `returvärde` deklareras sist i funktions signaturen
1. flera returvärden skrivs inom parenteser `()`
1. returnering avslutar en funktion

## Uppgift

1. Kan du summera negativa heltal?
1. Vad händer om du försöker summera flyttal, t.ex. 1.5 med 42?
1. Lägg till en funktion som räknar ut produkten av a och b, skriv ut svaret

[Exempel 10](../exempel9/README.md#exempel-10)
