package main

import "flag"
import "strings"
import "bufio"
import "os"
import "io"
import "strconv"

func konverti(de string, iksen bool) string {
	a := 0
	b := 1
	if iksen {
		a, b = 1, 0
	}
	tabelo := [18][2]string{
		{"cx", "ĉ"}, {"gx", "ĝ"}, {"hx", "ĥ"},
		{"jx", "ĵ"}, {"sx", "ŝ"}, {"ux", "ŭ"},

		{"Cx", "Ĉ"}, {"Gx", "Ĝ"}, {"Hx", "Ĥ"},
		{"Jx", "Ĵ"}, {"Sx", "Ŝ"}, {"Ux", "Ŭ"},

		{"CX", "Ĉ"}, {"GX", "Ĝ"}, {"HX", "Ĥ"},
		{"JX", "Ĵ"}, {"SX", "Ŝ"}, {"UX", "Ŭ"},
	}
	for i := 0; i<18; i++ {
		de = strings.Replace(de, tabelo[i][a], tabelo[i][b], -1)
	}
	return de
}

func konvertifluon(fluo *bufio.Reader, kien io.Writer, iksen bool) {
	for i, err:=fluo.ReadString('\n'); err == nil; i, err=fluo.ReadString('\n'){
		io.WriteString(kien, konverti(i, iksen))
	}
}

func main() {
	direkto := flag.Bool("x", false, "Traduki al iksoj. Convert to x-system.)")
	mimem := flag.Bool("i", false, "Skribi al la dosiero mem. In-place conversion of input file.")
	flag.Parse()

	enigo := flag.Arg(0)

	dosiero, err := os.Open(enigo)
	if err != nil {
		dosiero = os.Stdin
	}

	var kien io.ReadWriter
	if *mimem {
		kien, err = os.Create("/tmp/iksoj" + strconv.Itoa(os.Getpid()))
		if err != nil {panic("Oops, tempfile already exists!")}
	} else {
		kien = os.Stdout
	}
	konvertifluon(bufio.NewReader(dosiero), kien, *direkto)
	if *mimem {
		bufio.NewReader(kien).WriteTo(dosiero)
	}
}
