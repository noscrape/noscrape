package noscrape

import (
	"bytes"
	b64 "encoding/base64"
	"fmt"
	"seehuhn.de/go/postscript/funit"
	"seehuhn.de/go/postscript/type1"
	"seehuhn.de/go/sfnt"
	"seehuhn.de/go/sfnt/cff"
	"seehuhn.de/go/sfnt/cmap"
	"seehuhn.de/go/sfnt/glyph"
	"time"
)

func Render(font sfnt.Font, translation map[string]int32) (string, error) {

	notdefGlyph := cff.NewGlyph(".notdef", 1000) // Adjust the width as needed
	notdefGlyph.MoveTo(100, 100)
	notdefGlyph.LineTo(200, 100)
	notdefGlyph.LineTo(200, 200)
	notdefGlyph.LineTo(100, 200)

	cmapNew := cmap.Format12{}
	cmapNew[0xF000] = glyph.ID(0)

	encoding := make([]glyph.ID, len(translation)+1)
	encoding[0] = glyph.ID(0)

	b := []funit.Int16{0, 0, 0, 0}
	// origOutlines := n.font.Outlines.(*glyf.Outlines)
	var glyphs []*cff.Glyph

	glyphs = append(glyphs, notdefGlyph)

	for o, target := range translation {
		origin := []rune(o)[0]
		b = blueValues(font, origin, b)
		cmapNew[uint32(target)] = glyph.ID(len(glyphs))
		encoding = append(encoding, glyph.ID(len(glyphs)))
		origGlyf := fetchOriginalGlyph(font, origin)
		glyphs = append(glyphs, origGlyf)
	}

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

	newFont.InstallCMap(cmapNew)

	buf := new(bytes.Buffer)
	_, err := newFont.Write(buf)
	if err != nil {
		return "", fmt.Errorf("could not write tmp file: %v", err)
	}

	return b64.StdEncoding.EncodeToString(buf.Bytes()), nil
}
