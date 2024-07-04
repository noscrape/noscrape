package noscrape

import (
	"bytes"
	"fmt"
	"seehuhn.de/go/postscript/funit"
	"seehuhn.de/go/postscript/type1"
	"seehuhn.de/go/sfnt"
	"seehuhn.de/go/sfnt/cff"
	"seehuhn.de/go/sfnt/cmap"
	"seehuhn.de/go/sfnt/glyph"
	"time"
)

// Render generates a new obfuscated font based on the provided font and translation map.
// The translation map contains rune to glyph ID mappings used for obfuscation.
// It returns a buffer containing the new font data and an error if any occurred during the process.
func Render(font sfnt.Font, translation map[string]int32) (*bytes.Buffer, error) {

	// Create a .notdef glyph as a placeholder for undefined characters
	notdefGlyph := cff.NewGlyph(".notdef", 1000)
	notdefGlyph.MoveTo(100, 100)
	notdefGlyph.LineTo(200, 100)
	notdefGlyph.LineTo(200, 200)
	notdefGlyph.LineTo(100, 200)

	// Initialize a new cmap table
	cmapNew := cmap.Format12{}
	cmapNew[0xF000] = glyph.ID(0)

	// Initialize encoding and glyphs slices
	encoding := make([]glyph.ID, len(translation)+1)
	encoding[0] = glyph.ID(0)

	// Blue values used for the private dictionary in CFF outlines
	b := []funit.Int16{0, 0, 0, 0}
	var glyphs []*cff.Glyph

	// Add the .notdef glyph to the glyphs slice
	glyphs = append(glyphs, notdefGlyph)

	// Process each entry in the translation map
	for o, target := range translation {
		origin := []rune(o)[0]
		b = blueValues(font, origin, b)
		cmapNew[uint32(target)] = glyph.ID(len(glyphs))
		encoding = append(encoding, glyph.ID(len(glyphs)))
		origGlyf := fetchOriginalGlyph(font, origin)
		glyphs = append(glyphs, origGlyf)
	}

	// Create new CFF outlines with the updated glyphs and encoding
	newOutlines := &cff.Outlines{
		Private: []*type1.PrivateDict{
			{
				BlueValues: b,
			},
		},
		FDSelect: func(glyph.ID) int {
			return 0
		},
		Encoding: encoding,
		Glyphs:   glyphs,
	}

	// Create a new font with the obfuscated glyphs and updated metadata
	newFont := sfnt.Font{
		FamilyName: font.FamilyName + " (obfuscated by noscrape)",
		Width:      font.Width,
		Weight:     font.Weight,
		IsRegular:  font.IsRegular,

		CodePageRange: 0xF000,

		Version:          font.Version,
		CreationTime:     font.CreationTime,
		ModificationTime: time.Now(),

		UnitsPerEm: font.UnitsPerEm,
		FontMatrix: font.FontMatrix,

		Ascent:    font.Ascent,
		Descent:   font.Descent,
		LineGap:   font.LineGap,
		CapHeight: font.CapHeight,
		XHeight:   font.XHeight,

		UnderlinePosition:  font.UnderlinePosition,
		UnderlineThickness: font.UnderlineThickness,

		Outlines: newOutlines,
	}

	// Install the new cmap in the new font
	newFont.InstallCMap(cmapNew)

	// Write the new font data to a buffer
	buf := new(bytes.Buffer)
	_, err := newFont.Write(buf)
	if err != nil {
		return nil, fmt.Errorf("could not write tmp file: %v", err)
	}

	return buf, nil
}
