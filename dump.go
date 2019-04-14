// +build ignore

package main

import (
	"encoding/json"
	"io"
	"log"
	"os"

	mavlink "github.com/daedaleanai/gomavlink"
	"github.com/daedaleanai/gomavlink/ardupilotmega"
)

func main() {

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "\t")

	dec := mavlink.NewDecoder(os.Stdin, ardupilotmega.New)
	for {
		msg, strid, err := dec.Decode()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
		}
		if err == mavlink.ErrMustSync {
			n, err := dec.Resync()
			log.Println("resyncing: ", n, err)
		}

		log.Printf("%v: %T", strid, msg)
		if msg == nil {
			continue
		}
		if err := enc.Encode(msg); err != nil {
			log.Fatal(err)
		}
	}

}
