package highlight

// @version 0.0.2
// @license.name last updated at 2023/3/16 00:22:08

import (
	"fmt"
	"strconv"
)

type font int

const (
	standard font = iota
	bold
	Light
	Italic
	Underline
)

func Red(str string) string {
	return dye(0, "red", str)
}
func Green(str string) string {
	return dye(0, "green", str)
}
func Yellow(str string) string {
	return dye(0, "yellow", str)
}
func Blue(str string) string {
	return dye(0, "blue", str)
}
func Magenta(str string) string {
	return dye(0, "magenta", str)
}
func Cyan(str string) string {
	return dye(0, "cyan", str)
}
func White(str string) string {
	return dye(0, "white", str)
}
func Black(str string) string {
	return dye(0, "black", str)
}
func RED(str string) string {
	return dye(1, "red", str)
}
func GREEN(str string) string {
	return dye(1, "green", str)
}
func YELLOW(str string) string {
	return dye(1, "yellow", str)
}
func BLUE(str string) string {
	return dye(1, "blue", str)
}
func MAGENTA(str string) string {
	return dye(1, "magenta", str)
}
func CYAN(str string) string {
	return dye(1, "cyan", str)
}
func WHITE(str string) string {
	return dye(1, "white", str)
}
func BLACK(str string) string {
	return dye(1, "black", str)
}
func Redf(format string, args ...interface{}) string {
	return dye(2, "red", fmt.Sprintf(format, args...))
}
func Greenf(format string, args ...interface{}) string {
	return dye(2, "green", fmt.Sprintf(format, args...))
}
func Yellowf(format string, args ...interface{}) string {
	return dye(2, "yellow", fmt.Sprintf(format, args...))
}
func Bluef(format string, args ...interface{}) string {
	return dye(2, "blue", fmt.Sprintf(format, args...))
}
func Magentaf(format string, args ...interface{}) string {
	return dye(2, "magenta", fmt.Sprintf(format, args...))
}
func Cyanf(format string, args ...interface{}) string {
	return dye(2, "cyan", fmt.Sprintf(format, args...))
}
func Whitef(format string, args ...interface{}) string {
	return dye(2, "white", fmt.Sprintf(format, args...))
}
func Blackf(format string, args ...interface{}) string {
	return dye(2, "black", fmt.Sprintf(format, args...))
}
func REDf(format string, args ...interface{}) string {
	return dye(3, "red", fmt.Sprintf(format, args...))
}
func GREENf(format string, args ...interface{}) string {
	return dye(3, "green", fmt.Sprintf(format, args...))
}
func YELLOWf(format string, args ...interface{}) string {
	return dye(3, "yellow", fmt.Sprintf(format, args...))
}
func BLUEf(format string, args ...interface{}) string {
	return dye(3, "blue", fmt.Sprintf(format, args...))
}
func MAGENTAf(format string, args ...interface{}) string {
	return dye(3, "magenta", fmt.Sprintf(format, args...))
}
func CYANf(format string, args ...interface{}) string {
	return dye(3, "cyan", fmt.Sprintf(format, args...))
}
func WHITEf(format string, args ...interface{}) string {
	return dye(3, "white", fmt.Sprintf(format, args...))
}
func BLACKf(format string, args ...interface{}) string {
	return dye(3, "black", fmt.Sprintf(format, args...))
}

func dye(f font, color string, args ...interface{}) string {
	str := Concat(args...)
	n := "33"
	switch color {
	case "black":
		n = "30"
	case "red":
		n = "31"
	case "green":
		n = "32"
	case "yellow":
		n = "33"
	case "blue":
		n = "34"
	case "magenta":
		n = "35"
	case "cyan":
		n = "36"
	case "white":
		n = "37"
	}
	return fmt.Sprintf("\033[%s;%sm%s\033[0m", strconv.Itoa(int(f)), n, str)
}

func Concat(args ...interface{}) string {
	return join(" ", args...)
}

func join(separator string, args ...interface{}) (str string) {
	for i, param := range args {
		if i == len(args)-1 {
			separator = ""
		}
		str += fmt.Sprint(param) + separator
	}
	return
}
