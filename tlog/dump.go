// +build ignore

// Usage:
//    go run tlog/dump.go < ~/Downloads/flight.tlog 2>&1 | less
package main

import (
	"fmt"
	"io"
	"log"
	"os"

	tlog "../tlog"
)

func main() {

	dec := tlog.NewDecoder(os.Stdin)
	for {
		rec, err := dec.Decode()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
		}
		fmt.Println(&rec)
	}

}
