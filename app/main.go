package main

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/codecrafters-io/redis-starter-go/internal/resp"
)

var Store = make(map[any]any)

func Put(key, val any) error {
	Store[key] = val
	return nil
}

func Get(key any) (any, error) {
	val, ok := Store[key]
	if !ok {
		return nil, errors.New("couldn't get the key")
	}
	return val, nil
}

// Ensures gofmt doesn't remove the "net" and "os" imports in stage 1 (feel free to remove this!)
var _ = net.Listen
var _ = os.Exit

// func handlepongs(conn net.Conn) {
// 	writer := bufio.NewWriter(conn)
// 	writer.WriteString("+PONG\r\n")
// 	writer.Flush()
// }

func handleWrite(writer bufio.Writer, s string) {
	fmt.Println("entering writer handler for string: ", s)
	writer.WriteString(s)
	writer.Flush()
}

func toStringSlice(elements []any) []string {
	res := make([]string, len(elements))
	for i, val := range elements {
		switch v := val.(type) {
		case []byte:
			res[i] = string(v)
		default:
			res[i] = fmt.Sprintf("%v", elements[i])
		}
	}
	return res
}

//TODO: make a resp decoder?;
//func decodeRESP()

// we get it like *number\r\n$number\r\nPING\r\n -> we need to get these values before the other one
// *number -> number of words in the command like ping has 1, "echo hey" has two
// this needs to be appended
func handle(conn net.Conn) error {
	//fmt.Println("Entering handle")
	defer conn.Close()
	Reader := bufio.NewReader(conn)
	Writer := bufio.NewWriter(conn)
	for {
		//this doesnt have nested array capabilities yet
		elementsAny, err := resp.ParseRESP(Reader)
		if err != nil {
			fmt.Printf("encountered err : %v", err)
			return err
		}
		elements := toStringSlice(elementsAny)
		cmd := strings.ToUpper(elements[0])
		fmt.Println("entering switch for :", cmd)
		switch cmd {
		case "PING":
			handleWrite(*Writer, "+PONG\r\n")
		case "SET":
			if len(elements) != 3 {
				fmt.Printf("invalid usage of 'get'")
				return fmt.Errorf("invalid number of args with SET")
			}
			if err := Put(elements[1], elements[2]); err != nil {
				return err
			}
			handleWrite(*Writer, "+OK\r\n")
		case "GET":
			if len(elements) != 2 {
				return fmt.Errorf("invalid number of args with 'set'")
			}
			val, err := Get(elements[1])
			if err != nil {
				return err
			}
			returnString := "+" + val.(string) + "\r\n"
			handleWrite(*Writer, returnString)
		}
	}
}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment the code below to pass the first stage
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		go handle(conn)
	}

}
