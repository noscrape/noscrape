package main

import (
	"encoding/json"
	"fmt"
	"github.com/noscrape/noscrape/noscrape"
	"seehuhn.de/go/sfnt"
	"testing"
)

func TestHelloName(t *testing.T) {
	var m []noscrape.RuneMap
	s := noscrape.Obfuscate("This is a test.!?", m)

	j, _ := json.Marshal(s)

	fmt.Printf("%s\n", j)
	font, err := sfnt.ReadFile("./example/example.ttf")
	if err != nil {
		t.Fatalf("could not read font: %v", err)
	}

	var res noscrape.ObfuscationResult
	err = json.Unmarshal(j, &res)
	if err != nil {
		t.Fatalf("could not unmarshal json: %v", err)
	}

	r, e := noscrape.Render(*font, res.Map)
	if e != nil || r == "" {
		t.Fatalf("could not render")
	}

}
