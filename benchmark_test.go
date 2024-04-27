package main

import (
	"github.com/schoenbergerb/noscrape/noscrape"
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

	noscrape.Obfuscate("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz", m)
}
