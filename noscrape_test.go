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
	font, err := sfnt.ReadFile("/Users/bernhards/GolandProjects/noscrape/example/ubuntu.ttf")
	if err != nil {
		panic(err)
	}

	var res noscrape.ObfuscationResult
	err = json.Unmarshal(j, &res)
	if err != nil {
		panic(err)
	}

	r := noscrape.Render(*font, res.Map)
	println(r)
}
