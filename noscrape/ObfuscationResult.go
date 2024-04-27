package noscrape

type ObfuscationResult struct {
	Text string    `json:"text"`
	Map  []RuneMap `json:"map"`
}
