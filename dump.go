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

	dec := mavlink.NewDecoder(os.Stdin, ardupilotmega.Dialect)
	for {
		msg, sysid, compid, err := dec.Decode()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
		}
		log.Printf("%T sys:%d component:%d", msg, sysid, compid)
		if err := enc.Encode(msg); err != nil {
			log.Fatal(err)
		}
	}

}
