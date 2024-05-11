package main

/*
#include <stdlib.h>
*/
import "C"

import (
	"encoding/json"
	"github.com/noscrape/noscrape/noscrape"
	"seehuhn.de/go/sfnt"
)

//export noscrape_obfuscate
func noscrape_obfuscate(s *C.char, m *C.char) *C.char {
	var mapping []noscrape.RuneMap
	if m != nil {
		json.Unmarshal([]byte(C.GoString(m)), &mapping)
	}

	r := noscrape.Obfuscate(C.GoString(s), mapping)
	jsonData, err := json.Marshal(r)

	if err != nil {
		panic(err)
	}

	response := string(jsonData)

	return C.CString(response)
}

//export noscrape_render
func noscrape_render(f *C.char, m *C.char) *C.char {
	font, err := sfnt.ReadFile(C.GoString(f))
	if err != nil {
		panic(err)
	}

	var mapping []noscrape.RuneMap
	if m != nil {
		err = json.Unmarshal([]byte(C.GoString(m)), &mapping)
		if err != nil {
			panic(err)
		}
	}

	result := C.CString(noscrape.Render(*font, mapping))

	return result
}

func main() {
	// This is needed to build the shared library
}
