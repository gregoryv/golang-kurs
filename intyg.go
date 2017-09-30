package main

import "strings"
import "io/ioutil"
import "fmt"

func main() {
	var students = []string{
		"Orvar Otterstedt",
		"Ebba Otterstedt",
		"Edward Ericsson",
		"Arvid Ivarsson",
		"Simon Borginger",
		"Maria Borginger",
	}

	var tpl = `Närvarointyg för NAME

FIRST har närvarat vid __ av 3 tillfällen för introduktions
kurs i programmering. Kursen omfattade

+ Delar av GO programspråkets syntax
+ Grundläggande paket och funktioner
+ Enklare logiska konstruktioner

# Utbildare: Gregory Vincic
# Period: Från 2017-09-29 till 2017-10-01
`
	for _, name := range students {
		var filename = fmt.Sprintf("%s.txt", strings.Replace(name, " ", "_", 1))
		var names = strings.Split(name, " ")
		var tpl2 = strings.Replace(tpl, "NAME", name, 1)
		ioutil.WriteFile(
			filename,
			[]byte(strings.Replace(tpl2, "FIRST", names[0], 1)),
			0666,
		)
		print(filename, "\n")
	}
}
