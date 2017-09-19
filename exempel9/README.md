# Exempel 9

Funktioner med returvärden.

## main.go

    package main
	
	import "fmt"
	
	func sum(a, b int) int {
	    return a + b
	}
	
	func main() {
	    answer := sum(1,2)
	    fmt.Println(answer)
	}
	
Kom ihåg att spara filen innan du fortsätter. Kör programmet genom att i terminalen skriva

	go run main.go
	
## Lärdom

1. `returvärde` deklareras sist i funktions signaturen
2. `:=` kort konstruktion för att deklarera och tilldela en variabel ett värde

## Uppgift

1. Kan du summera negativa heltal?
1. Vad händer om du försöker summera flyttal, t.ex. 1.5 med 42?
1. Lägg till en funktion som räknar ut produkten av a och b, skriv ut svaret
