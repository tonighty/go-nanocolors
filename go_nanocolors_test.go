package go_nanocolors

import "testing"

func TestSimple(t *testing.T) {
	if !IsColorSupported() {
		t.Error("Color Is Not Supported")
	}

	if Red("text") != "\x1B[31mtext\x1B[39m" {
		t.Error("Not right")
	}
}

func TestVisual(t *testing.T) {
	println(Bold("Bold"))
	println(Dim("Dim"))
	println(Italic("Italic"))
	println(Underline("Underline"))
	println(Inverse("Inverse"))
	println(Hidden("Hidden"))
	println(Strikethrough("Strikethrough"))
	println(Black("Black"))
	println(Green("Green"))
	println(Yellow("Yellow"))
	println(Blue("Blue"))
	println(Magenta("Magenta"))
	println(Cyan("Cyan"))
	println(White("White"))
	println(Gray("Gray"))
	println(BgBlack("BgBlack"))
	println(BgRed("BgRed"))
	println(BgGreen("BgGreen"))
	println(BgYellow("BgYellow"))
	println(BgBlue("BgBlue"))
	println(BgMagenta("BgMagenta"))
	println(BgCyan("BgCyan"))
	println(BgWhite("BgWhite"))

	println(RGB(250, 250, 250)("kek"))
}
