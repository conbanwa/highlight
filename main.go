package highlight

import (
	"fmt"
	"strconv"
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

func dye(highlight int, color string, args ...interface{}) string {
	str := Concat(args...)
	n := "37"
	switch color {
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
	default:
		n = "33"
	}
	return fmt.Sprintf("\033[%s;%sm%s\033[0m", strconv.Itoa(highlight), n, str)
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
