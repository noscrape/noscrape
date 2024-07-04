package noscrape

// RuneMap represents a mapping between an original rune and its obfuscation target.
// It is used to obfuscate text by replacing original runes with their corresponding targets.
type RuneMap struct {
	// OriginalRune is the original rune that will be obfuscated.
	OriginalRune rune `json:"origin"`

	// ObfuscationTarget is the target value to replace the original rune with.
	// It is represented as an int32.
	ObfuscationTarget int32 `json:"target"`
}
