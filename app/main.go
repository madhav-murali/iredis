package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"regexp"
	"strconv"
)

// Ensures gofmt doesn't remove the "net" and "os" imports in stage 1 (feel free to remove this!)
var _ = net.Listen
var _ = os.Exit

func handlepongs(conn net.Conn) {
	writer := bufio.NewWriter(conn)
	writer.WriteString("+PONG\r\n")
	writer.Flush()
}

// we get it like *number\r\n$number\r\nPING\r\n -> we need to get these values before the other one

func handle(conn net.Conn) {
	//fmt.Println("Entering handle")
	defer conn.Close()

	reader := bufio.NewReader(conn)
	//writer := bufio.NewWriter(conn)

	for {
		//msg := scanner.Text()
		//log.Println("received msg :", msg)
		numCmds, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("error reading number of commands:", err)
		}
		re := regexp.MustCompile(`\d+`)
		matches := re.FindAllString(numCmds, -1)
		numCmds_, err := strconv.Atoi(matches[0])
		for range numCmds_ {
			_, err := reader.ReadString('\n') //now we need to read the $ number
			if err != nil {
				fmt.Println("error reading from client")
				return
			}
			_, err = reader.ReadString('\n')
			handlepongs(conn)
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
	fmt.Println("Why is this happening")
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		go handle(conn)
	}

}
