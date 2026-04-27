package storage

import (
	"errors"
	"log"
	"strconv"
	"time"
)

func HandleSet(c *Cache, elements []string) error {
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

// func HandleGet(c *Cache, elements []string) (any, error) {

// 	// if len(elements) != 2 {
// 	// 	return fmt.Errorf("invalid number of args with 'set'")
// 	// }
// 	// val, err := Get(elements[1])
// 	// if err != nil {
// 	// 	return err
// 	// }
// 	// valString := val.(string)
// 	// //returnString := "$" + strconv.Itoa(len(valString)) + "\r\n" + valString + "\r\n"
// 	// handleWrite(*Writer, resp.EchoRESP(valString))

// 	return c.Get(elements[1])
// }
