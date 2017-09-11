# Exempel 1

Vad tror du nedan program gör?

Skriv programmet i en textfil och döp denna till `main.go`

## main.go

	package main

	func main() {
		print("Hello, World!")
	}

Kör programmet genom att i terminalen skriva

    go run main.go
	

## Lärdom

- `nyckelord` ett programeringspråk består av nyckelord med olika betydelse, exemplet ovan har två nyckelord, `package` och `func`
- `deklarerar` skapar, eller definierar en betydelse i programmet
- `package` nyckelord som deklarerar vilket paket denna koden tillhör
- `func` nyckelord som deklarerar en funktion
- `terminalen`, programmet som vi använder för att skicka kommandon till datorn
- `källkoden` som skrivs i en textfil, vi döpte denna till `main.go`
- kompilatorn `go`, verktyget som omvandlar källkoden till något som datorn förstår
- `main` huvud funktionen i varje program om den ligger i paketet `main`
- funktionsanrop till `print`
- funktioner kan ta ett eller flera argument `"Hello, World!"`
- en `sträng` måste skrivas med dubbelfnutt `"..."` i början och slut


## Datorn

Ditt program pratar med operativsystemet som i sin tur använder drivrutiner för 
att prata med olika hårdvarukomponenter.

![Översikt](overview.png)
