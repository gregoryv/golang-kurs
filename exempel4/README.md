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
	
Kom ihåg att spara filen innan du fortsätter. Kör programmet genom att i terminalen skriva

	go run main.go
	
	
## Lärdom

- `variabel` är en plats i minnet som kan ha ett värde av en viss typ
- `var` nyckelordet används för att deklarera en ny variabel
- man kan även hoppa över var deklarationen om man direkt sätter ett värde, då använder man `:=` 

## Uppgift

- Lägg till en variabel `c` med värde `9`
- Skriv ut summan av `a+b+c`
