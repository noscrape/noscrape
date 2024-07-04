package noscrape

// ObfuscationResult holds the result of an obfuscation operation.
// It contains the obfuscated text and the map of rune obfuscations applied.
type ObfuscationResult struct {
	// Text is the obfuscated version of the input text.
	Text string `json:"text"`

	// Map is a slice of RuneMap structs that describe the obfuscations applied to the original text.
	Map []RuneMap `json:"map"`
}
