package resp

import (
	"strconv"
	"strings"
)

func EchoRESP(input string) string {
	return "$" + strconv.Itoa(len(input)) + "\r\n" + input + "\r\n"
}
func RESPstring(input []string) string {
	var s strings.Builder
	s.WriteString("*" + strconv.Itoa(len(input)) + "\r\n")
	for _, val := range input {
		s.WriteString("$" + strconv.Itoa(len(val)) + "\r\n")
		s.WriteString(val + "\r\n")
	}
	return s.String()
}
