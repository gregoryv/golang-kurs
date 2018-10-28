# Exempel 4

Variabler och operatorer.

## main.go

	package main

	import "fmt"

	func main() {
		var a int = 5
		b := 7
		fmt.Println(a + b)
	}

Kom ihåg att spara filen innan du fortsätter. Kör programmet genom att
i terminalen skriva

	go run main.go


## Lärdom

1. `variabel` är en plats i minnet som kan ha ett värde av en viss typ
1. `var` nyckelordet används för att deklarera en ny variabel
1. man kan även hoppa över var deklarationen om man direkt sätter ett
   värde, då använder man `:=`

## Uppgift

1. Lägg till en variabel `c` med värde `9`
1. Skriv ut summan av `a+b+c`

[Exempel 5](../05/README.md#exempel-5)
