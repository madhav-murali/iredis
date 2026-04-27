package storage

import (
	"bufio"
	"errors"
	"log"
	"strconv"
	"time"
)

func HandleSet(c *Cache, writer bufio.Writer, elements []string) error {
	// if len(elements) != 3 {
	// 	fmt.Printf("invalid usage of 'get'")
	// 	return fmt.Errorf("invalid number of args with SET")
	// }
	var ttl time.Duration
	if len(elements) == 3 {
		log.Printf("no limit set")
	} else if len(elements) == 5 {
		tm, err := strconv.Atoi(elements[4])
		if err != nil {
			return err
		}
		if elements[3] == "PX" {
			ttl = time.Duration(tm) * time.Millisecond
		} else {
			ttl = time.Duration(tm) * time.Second
		}
	} else {
		return errors.New("invalid usage of Set")
	}

	if err := c.Set(elements[1], elements[2], ttl); err != nil {
		return err
	}
	// handleWrite(*Writer, "+OK\r\n")
	return nil
}
