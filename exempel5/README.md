# Exempel 5

Felsökning, kompilatorn är din vän.

## main.go

	package main
	
	func main() {
		fmt.Println("Hello, World!")
	}
	
Kom ihåg att spara filen innan du fortsätter. Kör programmet genom att i terminalen skriva

	go run main.go
	
Vad hände? varför skrev det inte ut `Hello, World!`?

	./main.go:4: undefined: fmt in fmt.Println

## Lärdom

- Kompilatorn berättar var felet ligger dvs. vilken fil och rad samt att den säger vad som är fel

## Uppgift

- fixa felet
- gör ett medvetet fel i koden som att t.ex. ta bort ett av `"`, spara och försök köra programmet igen

[Exempel 6](../exempel6/README.md#exempel-6)
