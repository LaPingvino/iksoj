package main

import "fmt"
import "flag"
import "strings"

func konverti(de string, kien bool) string {
	returnstring := de
	a := 0
	b := 1
	if kien {
		a, b = 1, 0
	}
	tabelo := [6][2]string{
		{"cx", "ĉ"},
		{"gx", "ĝ"},
		{"hx", "ĥ"},
		{"jx", "ĵ"},
		{"sx", "ŝ"},
		{"ux", "ŭ"},
	}
	for i := 0; i<6; i++ {
		returnstring = strings.Replace(returnstring, tabelo[i][a], tabelo[i][b], -1)
	}
	return returnstring
}

func main() {
	var direkto *bool
	var enigo string

	direkto = flag.Bool("iksen", false, "Traduki al iksoj")

	flag.Parse()
	enigo = strings.Join(flag.Args()," ")
	fmt.Print(konverti(enigo,*direkto))
}
