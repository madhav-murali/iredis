package resp

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func ParseRESP(r *bufio.Reader) ([]any, error) {
	//first get the number of
	line, err := r.ReadString('\n')
	if err != nil {
		return nil, err
	}
	line = strings.TrimSpace(line)
	count, err := strconv.Atoi(line[1:])
	if err != nil {
		return nil, err
	}

	elements := make([]any, count)
	for i := range count {
		//TODO: change to a more generic read function
		lineLength, err := parseBulkLength(r)
		if err != nil {
			return nil, err
		}
		elem, err := readFullBody(r, lineLength)
		if err != nil {
			return nil, err
		}
		elements[i] = elem
	}
	return elements, nil
}

func parseBulkLength(r *bufio.Reader) (int, error) {
	//read the line like $5\r\n
	line, err := r.ReadString('\n')
	if err != nil {
		return 0, err
	}

	if line[0] != '$' || len(line) < 3 {
		return 0, fmt.Errorf("invalid bulk string prefix")
	}

	line = strings.TrimSpace(line)
	return strconv.Atoi(line[1:])
}

func readFullBody(r *bufio.Reader, length int) ([]byte, error) {
	buf := make([]byte, length)
	_, err := io.ReadFull(r, buf)
	if err != nil {
		return nil, err
	}
	_, err = r.Discard(2)
	return buf, err
}
