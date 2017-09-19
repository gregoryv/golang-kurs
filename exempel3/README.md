# Exempel 3

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

- `display` är en egen funktion som tar ett argument `something`
- `string` datatypen är en sträng, kompilatorn hjälper dig att skilja på olika datatyper.
- `datatyper` kan vara sträng, tecken, siffra, decimaltal osv.

[Exempel 4](../exempel4/README.md)
