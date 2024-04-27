package main

import (
	"C"
	"encoding/json"
	"seehuhn.de/go/sfnt"

	"github.com/schoenbergerb/noscrape/noscrape"
)

//export noscrape_obfuscate
func noscrape_obfuscate(s *C.char) *C.char {
	var m []noscrape.RuneMap
	r := noscrape.Obfuscate(C.GoString(s), m)
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

	var res noscrape.ObfuscationResult
	err = json.Unmarshal([]byte(C.GoString(m)), &res)
	if err != nil {
		panic(err)
	}

	return C.CString(noscrape.Render(*font, res.Map))
}

func main() {
	// This is needed to build the shared library
}
