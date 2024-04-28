package main

import (
	"github.com/schoenbergerb/noscrape/noscrape"
	"seehuhn.de/go/sfnt"
	"testing"
)

func BenchmarkObfuscation(b *testing.B) {
	var m []noscrape.RuneMap

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		noscrape.Obfuscate("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz", m)
	}
}

func BenchmarkRendering(b *testing.B) {

	var m []noscrape.RuneMap

	mt := noscrape.Obfuscate("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz", m)

	font, err := sfnt.ReadFile("./example/example.ttf")
	if err != nil {
		panic(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		noscrape.Render(*font, mt.Map)
	}
}
