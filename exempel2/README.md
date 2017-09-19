# Exempel 2

Ändra det första exemplet till följande.

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

- `import` används för att återanvända kod i andra `paket`
- `fmt.Println` funktionsanrop till andra paket
- publika funktioner i paket börjar med Versaler, privata med gemener
- vilka paket finns? det finns massor, https://godoc.org/-/go är en bra början
	
[Exempel 3](../exempel3/README.md#exempel-3)
