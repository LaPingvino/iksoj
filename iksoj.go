package main

import "flag"
import "strings"
import "bufio"
import "os"
import "io"

func konverti(de string, iksen bool) string {
	returnstring := de
	a := 0
	b := 1
	if iksen {
		a, b = 1, 0
	}
	tabelo := [18][2]string{
		{"cx", "ĉ"},
		{"gx", "ĝ"},
		{"hx", "ĥ"},
		{"jx", "ĵ"},
		{"sx", "ŝ"},
		{"ux", "ŭ"},
		{"Cx", "Ĉ"},
		{"Gx", "Ĝ"},
		{"Hx", "Ĥ"},
		{"Jx", "Ĵ"},
		{"Sx", "Ŝ"},
		{"Ux", "Ŭ"},
		{"CX", "Ĉ"},
		{"GX", "Ĝ"},
		{"HX", "Ĥ"},
		{"JX", "Ĵ"},
		{"SX", "Ŝ"},
		{"UX", "Ŭ"},
	}
	for i := 0; i<18; i++ {
		returnstring = strings.Replace(returnstring,
				tabelo[i][a], tabelo[i][b], -1)
	}
	return returnstring
}

func konvertifluon(fluo *bufio.Reader, kien io.Writer, iksen bool) {
	for i, err:=fluo.ReadString('\n'); err == nil; i, err=fluo.ReadString('\n'){
		io.WriteString(kien, konverti(i, iksen))
	}
}

func main() {
	direkto := flag.Bool("iksen", false, "Traduki al iksoj. Defaŭlte tradukiĝas al veraj ĉapeloj.")
	flag.Parse()

	enigo := flag.Arg(0)

	dosiero, err := os.Open(enigo)
	if err != nil {
		dosiero = os.Stdin
	}
	konvertifluon(bufio.NewReader(dosiero), os.Stdout, *direkto)
}
