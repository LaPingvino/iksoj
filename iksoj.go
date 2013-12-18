package main

import "fmt"
import "flag"

func konverti(de string, kien bool) string {
	return ""
}

func main() {
	var direkto *bool
	direkto = flag.Bool("iksen", false, "Traduki al iksoj")
	flag.Parse()
	fmt.Println(*direkto)
}
