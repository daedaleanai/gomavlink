// Gomavgen generates a Go package from a MAVLink dialect definition xml file and it's includes.
package main

import (
	"bytes"
	"encoding/xml"
	"flag"
	"go/format"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func main() {

	log.SetFlags(0)
	log.SetPrefix("gomavgen: ")
	flag.Parse()

	if len(flag.Args()) != 1 {
		log.Fatalf("Usage: %s path/to/dialect.xml")
	}
	f, err := os.Open(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	dname, fname := filepath.Split(f.Name())
	basename := strings.TrimSuffix(fname, filepath.Ext(fname))

	dialect := MAVLink{Name: basename}

	enums := map[string]*Enum{}

	if err := xml.NewDecoder(f).Decode(&dialect); err != nil {
		log.Fatal(err)
	}
	f.Close()

	// The spec says only includes in the top level xml are executed, probably because proper
	// recursion was beyond the grasp of the designers.
	// The spec mentions repeated enum definitions (see below), but leaves the semantics
	// of repeated enum entries or repeated message declarations open, so we'll naively process
	// and leave it to the Go compiler to flag redefinitions.
	for _, v := range dialect.Include {
		f, err := os.Open(filepath.Join(dname, v))
		if err != nil {
			log.Fatal(err)
		}
		var inc MAVLink
		if err := xml.NewDecoder(f).Decode(&inc); err != nil {
			log.Fatal(v, ":", err)
		}
		log.Printf("Including %s dialect %d version %d", f.Name(), inc.Dialect, inc.Version)
		f.Close()

		// Enum declarations are to be merged.
		for _, vv := range inc.Enums {
			if enums[vv.Name] == nil {
				enums[vv.Name] = vv
				dialect.Enums = append(dialect.Enums, vv)
				continue
			}
			log.Println("Merging %q enum %q", v, vv.Name)
			enums[vv.Name].Entries = append(enums[vv.Name].Entries, vv.Entries...)
		}

		dialect.Messages = append(dialect.Messages, inc.Messages...)

		if dialect.Version == 0 {
			log.Println("Inheriting version from", v)
			dialect.Version = inc.Version
		}
	}

	log.Printf("Generating package %s dialect %d version %d", basename, dialect.Dialect, dialect.Version)

	// stable reorder fields by their scalar size
	for _, v := range dialect.Messages {
		sort.Stable(bySerialisationOrder(v.Fields))
	}

	if err := os.MkdirAll(basename, 0755); err != nil {
		log.Fatalln(err)
	}

	of, err := os.Create(filepath.Join(basename, "mavlink.go"))
	if err != nil {
		log.Fatalln(err)
	}
	defer of.Close()

	var b bytes.Buffer
	if err := tmpl.Execute(&b, dialect); err != nil {
		log.Fatal(err)
	}
	out := b.Bytes()
	out, err = format.Source(out)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := of.Write(out); err != nil {
		log.Fatal(err)
	}

}
