package go_nanocolors

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strings"
)

var _, noColor = os.LookupEnv("NO_COLOR")

var isDisabled = noColor || *flag.Bool("--no-color", false, "disable colors")

var _, forceColor = os.LookupEnv("FORCE_COLOR")

var isForced = forceColor || *flag.Bool("--color", false, "force color")

var isWindows = runtime.GOOS == "windows"

// TODO: tty
var isCompatibleTerminal = os.Getenv("TERM") != "dumb"

var isEnabled = !isDisabled && (isForced || isWindows || isCompatibleTerminal)

func color(open, close, closeRegexp, restore string) func(s string) string {
	if isEnabled {
		return func(s string) string {
			if len(s) > 0 {
				return open + s + close
			} else {
				if strings.Index(s, close) > 4 {
					s = strings.ReplaceAll(s, closeRegexp, restore)
				}
				return open + s + close
			}
		}
	} else {
		return func(s string) string {
			return s
		}
	}
}

const close22 = "\x1b[22m"
const close39 = "\x1b[39m"
const close49 = "\x1b[49m"
const regexp22 = `/\x1b\[22m/g`
const regexp39 = `/\x1b\[39m/g`
const regexp49 = `/\x1b\[49m/g`

var Bold = color("\x1b[1m", close22, regexp22, "\x1b[22m\x1b[1m")
var Dim = color("\x1b[2m", close22, regexp22, "\x1b[22m\x1b[2m")
var Italic = color("\x1b[3m", "\x1b[23m", `/\x1b\[23m/g`, "\x1b[3m")
var Underline = color("\x1b[4m", "\x1b[24m", `/\x1b\[24m/g`, "\x1b[4m")
var Inverse = color("\x1b[7m", "\x1b[27m", `/\x1b\[27m/g`, "\x1b[7m")
var Hidden = color("\x1b[8m", "\x1b[28m", `/\x1b\[28m/g`, "\x1b[8m")
var Strikethrough = color("\x1b[9m", "\x1b[29m", `/\x1b\[29m/g`, "\x1b[9m")
var Black = color("\x1b[30m", close39, regexp39, "\x1b[30m")
var Red = color("\x1b[31m", close39, regexp39, "\x1b[31m")
var Green = color("\x1b[32m", close39, regexp39, "\x1b[32m")
var Yellow = color("\x1b[33m", close39, regexp39, "\x1b[33m")
var Blue = color("\x1b[34m", close39, regexp39, "\x1b[34m")
var Magenta = color("\x1b[35m", close39, regexp39, "\x1b[35m")
var Cyan = color("\x1b[36m", close39, regexp39, "\x1b[36m")
var White = color("\x1b[37m", close39, regexp39, "\x1b[37m")
var Gray = color("\x1b[90m", close39, regexp39, "\x1b[90m")
var BgBlack = color("\x1b[40m", close49, regexp49, "\x1b[40m")
var BgRed = color("\x1b[41m", close49, regexp49, "\x1b[41m")
var BgGreen = color("\x1b[42m", close49, regexp49, "\x1b[42m")
var BgYellow = color("\x1b[43m", close49, regexp49, "\x1b[43m")
var BgBlue = color("\x1b[44m", close49, regexp49, "\x1b[44m")
var BgMagenta = color("\x1b[45m", close49, regexp49, "\x1b[45m")
var BgCyan = color("\x1b[46m", close49, regexp49, "\x1b[46m")
var BgWhite = color("\x1b[47m", close49, regexp49, "\x1b[47m")

var RGB = func(red, green, blue uint8) func(s string) string {
	open := fmt.Sprintf("\u001B[38;2;%v;%v;%vm", red, green, blue)
	return color(open, close39, regexp39, open)
}

func IsColorSupported() bool {
	return isEnabled
}
