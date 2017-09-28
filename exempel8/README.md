# Exempel 8

Uttryck och operatorer

## main.go

	package main

	import "fmt"

	func main() {
        fmt.Print("1 + 1 + 5 = ", 1+1+5, "\n")
        fmt.Print("1 - 1 = ", 1-1, "\n")
    }

Kom ihåg att spara filen innan du fortsätter. Kör programmet genom att i terminalen skriva

	go run main.go

## Lärdom

1. `uttryck` kan vara beräkningar av ett eller flera tal
1. `operator` används vid beräkningar, t.ex. '+-*/%'

## Uppgift

1. Lägg till ett uttryck för varje operator
  - `*` multiplikation
  - `/` division
  - `%` modulus
1. Vad gör modulus operatorn?

[Exempel 9](../exempel9/README.md#exempel-9)
