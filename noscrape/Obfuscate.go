package noscrape

import (
	"math/rand"
)

/*
Obfuscate takes a string and returns an obfuscated version of the string.

Parameters:
  - n: A pointer to a Noscrape instance containing the translation map used for obfuscation.
  - s: The input string to be obfuscated.

Returns:
  - A string representing the obfuscated version of the input string.

Note:
  - The obfuscation process is reversible only if the original Noscrape instance used for obfuscation is available.
  - The obfuscated string may contain non-printable Unicode characters as identifiers.

Example:
n := &Noscrape{translation: make(map[rune]int32)}
obfuscated := Obfuscate(n, "Hello, World!")
fmt.Println(obfuscated) // Output may vary: "㋠㋃㋈㋂⧁㋞㋛⧂㋗㋟㋍"

go
Copy code

Reference:
- Fisher-Yates shuffle algorithm: https://en.wikipedia.org/wiki/Fisher%E2%80%93Yates_shuffle

See also:
- Noscrape struct
*/
func Obfuscate(s string, translation []RuneMap) ObfuscationResult {
	// shuffle runes
	inRune := []rune(s)
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})

	fetchMapping := func(r rune) (RuneMap, bool) {
		for _, mapping := range translation {
			if mapping.OriginalRune == r {
				return mapping, true
			}
		}
		return RuneMap{}, false
	}
	// calc custom unicode
	for _, r := range inRune {
		if _, exists := fetchMapping(r); !exists {
			translation = append(translation, RuneMap{
				OriginalRune:      r,
				ObfuscationTarget: int32(0xF000 + len(translation) + 1),
			})
		}
	}

	// translate provided string
	runes := make([]rune, 0)
	for _, r := range s {
		m, _ := fetchMapping(r)
		runes = append(runes, m.ObfuscationTarget)
	}

	return ObfuscationResult{
		Text: string(runes),
		Map:  translation,
	}
}
