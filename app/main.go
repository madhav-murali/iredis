package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/codecrafters-io/redis-starter-go/internal/resp"
	"github.com/codecrafters-io/redis-starter-go/internal/storage"
)

// var c storage.Cache

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

// we get it like *number\r\n$number\r\nPING\r\n -> we need to get these values before the other one
// *number -> number of words in the command like ping has 1, "echo hey" has two
// this needs to be appended
func handle(conn net.Conn, c *storage.Cache, lst *storage.List) error {
	//fmt.Println("Entering handle")
	defer conn.Close()
	Reader := bufio.NewReader(conn)
	Writer := bufio.NewWriter(conn)
	for {
		elementsAny, err := resp.ParseRESP(Reader)
		if err != nil {
			if err == io.EOF {
				fmt.Println("client closed the connection")
				return nil
			}
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
			if err := storage.HandleSet(c, elements); err != nil {
				handleWrite(*Writer, "$-1\r\n") // handling nil
			} else {
				handleWrite(*Writer, "+OK\r\n")
			}
		case "GET":
			val, ok := c.Get(elements[1])
			if !ok {
				handleWrite(*Writer, "$-1\r\n")
				return errors.New("invalid key or has expired")
			}
			valString := val.(string)
			//returnString := "$" + strconv.Itoa(len(valString)) + "\r\n" + valString + "\r\n"
			handleWrite(*Writer, resp.EchoRESP(valString))

		case "ECHO":
			writeString := strings.Join(elements[1:], "")
			handleWrite(*Writer, resp.EchoRESP(writeString))
		case "RPUSH":
			length := lst.RPUSH(elements[1], elements[2:])
			handleWrite(*Writer, ":"+strconv.Itoa(length)+"\r\n")
		case "LRANGE":
			handleWrite(*Writer, resp.RESPstring(lst.LRANGE(elements[1], elements[2], elements[3])))
		case "LPUSH":
			length := lst.LPUSH(elements[1], elements[2:])
			handleWrite(*Writer, ":"+strconv.Itoa(length)+"\r\n")
		case "LLEN":
			handleWrite(*Writer, ":"+strconv.Itoa(lst.LLEN(elements[1]))+"\r\n")
		}
	}
}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")
	c := storage.NewCache()
	list := storage.NewList()
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
		go handle(conn, c, list)
	}

}
