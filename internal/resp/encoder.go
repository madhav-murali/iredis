package resp

import "strconv"

func EchoRESP(input string) string {
	return "$" + strconv.Itoa(len(input)) + "\r\n" + input + "\r\n"
}
