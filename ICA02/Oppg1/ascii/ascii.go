package ascii

import "fmt"

const ascii = "\x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\x0a\x0b\x0c\x0d\x0e\x0f" +
	"\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1c\x1d\x1e\x1f" +
	` !"#$%&'()*+,-./0123456789:;<=>?` +
	`@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_` +
	"`abcdefghijklmnopqrstuvwxyz{|}~\x7f"

// Oppgave 1b
// Implementer en funksjon som eksportere const ascii

// Funksjon tar en "string literal" med kun ASCII tegn og lager en utskrift på
// følgende format:
// [ascii-kode heksadesimalt med store bokstaver A-F][mellomrom]
// [symbol for ascii-kode][mellomrom][ascii-kode binært][linjeskift]
//
// Eksempel (utskriften kommer fra en main.go fil):
//	…
// 3E > 111110
// 3F ? 111111
// 40 @ 1000000
func IterateOverASCIIStringLiteral() {
	for i := 0; i < len(ascii); i++ {
		fmt.Printf("%X %q %b\n", ascii[i], ascii[i], ascii[i])
	}
}

func GetASCIIStringLiteral() string {
	return ascii
}

func GreetingASCII() {
	fmt.Printf("Hello :-)")
}
