# Exempel 7

Nyckelorden `if` och `else`.



## main.go

	package main
	
	import "fmt"
	
	func main() {
        fmt.Print("Hello ")
        if true {
            fmt.Println("Kalle!")
        } else {
            fmt.Println("Lisa!")
        }
    }
	
Kom ihåg att spara filen innan du fortsätter. Kör programmet genom att i terminalen skriva

	go run main.go
	
Vad tror du kommer hända?

## Lärdom

1. `if` börjar ett uttryck som skall utvärderas till sant eller falskt
1. `else` kan användas som en alternativ väg
1. `true` och `false` är boolska värden

## Uppgift

1. Ändra uttrycket i `if`-satsen till falskt, vad kommer skrivas ut?

[Exempel 8](../08/README.md#exempel-8)
