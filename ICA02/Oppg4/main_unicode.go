package main

import (
	"github.com/Snaxai/IS105/ICA02/Oppg4/unicode"
)

func main() {
	unicode.Gethexadesimal("“北と南”")
	unicode.Gethexadesimal("“norður og suður”")
	unicode.Gethexadesimal("“nord og sør”")
	unicode.Translate("nord og sør", "is")
	unicode.Translate("nord og sør", "jp")
	unicode.UnicodeCodePointDemo()
}
