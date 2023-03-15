package highlight

import (
	"fmt"
	"strconv"
)

func Red(str string) string {
	return dye(0, "31", str)
}
func Green(str string) string {
	return dye(0, "32", str)
}
func Yellow(str string) string {
	return dye(0, "33", str)
}
func Blue(str string) string {
	return dye(0, "34", str)
}
func Magenta(str string) string {
	return dye(0, "35", str)
}
func Cyan(str string) string {
	return dye(0, "36", str)
}
func White(str string) string {
	return dye(0, "37", str)
}
func Black(str string) string {
	return dye(0, "30", str)
}
func RED(str string) string {
	return dye(1, "31", str)
}
func GRE(str string) string {
	return dye(1, "32", str)
}
func YEL(str string) string {
	return dye(1, "33", str)
}
func BLU(str string) string {
	return dye(1, "34", str)
}
func MAG(str string) string {
	return dye(1, "35", str)
}
func CYA(str string) string {
	return dye(1, "36", str)
}
func WHI(str string) string {
	return dye(1, "37", str)
}
func BLA(str string) string {
	return dye(1, "30", str)
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
	default:
		n = "33"
	}
	return fmt.Sprintf("\033["+strconv.Itoa(highlight)+";"+n+";40m%s\033[0m", str)
}

func Concat(args ...interface{}) string {
	return Join(" ", args...)
}

func Join(separator string, args ...interface{}) (str string) {
	for i, param := range args {
		if i == len(args)-1 {
			separator = ""
		}
		str += fmt.Sprint(param) + separator
	}
	return
}
