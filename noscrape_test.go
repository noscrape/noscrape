package main

import (
	"github.com/noscrape/noscrape/noscrape"
	"seehuhn.de/go/sfnt"
	"testing"
)

func TestSimple(t *testing.T) {

}

func TestHelloName(t *testing.T) {
	m := map[string]int32{
		"a": 0xebda,
	}

	font, err := sfnt.ReadFile("./example/example.ttf")
	if err != nil {
		t.Fatalf("could not read font: %v", err)
	}

	r, e := noscrape.Render(*font, m)
	if e != nil || r == "" {
		t.Fatalf("could not render")
	}

	println(r)
}
