// Gomavgen generates a Go package from a MAVLink dialect definition xml file and its includes.
package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"text/template"
)

func main() {

	log.SetFlags(0)
	log.SetPrefix("gomavgen: ")
	flag.Parse()

	if len(flag.Args()) != 2 {
		log.Fatalf("Usage: %s path/to/lang.tmpl path/to/dialect.xml", os.Args[0])
	}

	tmpl, err := template.New(filepath.Base(flag.Arg(0))).Funcs(tmplfuncs).ParseFiles(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("template file:", tmpl.Name())

	f, err := os.Open(flag.Arg(1))
	if err != nil {
		log.Fatal(err)
	}
	dname, fname := filepath.Split(f.Name())
	basename := strings.ToLower(strings.TrimSuffix(fname, filepath.Ext(fname)))

	dialect := MAVLink{Name: basename}

	enums := map[string]*Enum{}

	if err := xml.NewDecoder(f).Decode(&dialect); err != nil {
		log.Fatal(err)
	}
	log.Printf("Top level %s dialect %d version %d", f.Name(), dialect.Dialect, dialect.Version)
	f.Close()

	for _, v := range dialect.Enums {
		enums[v.Name] = v
	}

	// The spec says only includes in the top level xml are executed.
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

		// Enum declarations are to be merged, where the includes are supposed to come first
		for _, vv := range inc.Enums {
			if enums[vv.Name] == nil {
				enums[vv.Name] = vv
				dialect.Enums = append(dialect.Enums, vv)
				continue
			}
			log.Printf("Merging %q enum %q", v, vv.Name)
			enums[vv.Name].Entries = append(vv.Entries, enums[vv.Name].Entries...)
		}

		dialect.Messages = append(dialect.Messages, inc.Messages...)

		if dialect.Version == 0 {
			log.Println("Inheriting version from", v)
			dialect.Version = inc.Version
		}
	}

	log.Printf("Generating package %s dialect %d version %d", basename, dialect.Dialect, dialect.Version)

	// fill in missing enum values, starting from highest found (?)
	for _, v := range dialect.Enums {
		max := uint64(0)
		for _, vv := range v.Entries {
			if vv.Value != "" {
				val, _ := strconv.ParseUint(vv.Value, 0, 32)
				if max < val {
					max = val
				}
			}
		}
		warnenums := false
		for i, vv := range v.Entries {
			if vv.Value == "" {
				if uint64(i) != max+1 {
					warnenums = true
				}
				vv.Value = fmt.Sprintf("%d", max+1)
				max++
			}
		}
		if warnenums {
			log.Printf("Possibly ill-defined mixing of explicit and implicit values in enum %s may be inconsistent", v.Name)
		}
	}

	sort.Sort(byMessageID(dialect.Messages))
	// stable reorder fields by their scalar size
	for _, v := range dialect.Messages {
		sort.Stable(bySerialisationOrder(v.Fields))
	}

	if err := tmpl.Execute(os.Stdout, dialect); err != nil {
		log.Fatal(err)
	}

}
