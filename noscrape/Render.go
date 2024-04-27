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

func Render(font sfnt.Font, translation []RuneMap) string {

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

	for _, m := range translation {
		fmt.Printf("%v ", m.OriginalRune)
		b = blueValues(font, m.OriginalRune, b)
		cmapNew[uint32(m.ObfuscationTarget)] = glyph.ID(len(glyphs))
		encoding = append(encoding, glyph.ID(len(glyphs)))
		origGlyf := fetchOriginalGlyph(font, m.OriginalRune)
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
	_, err1 := newFont.Write(buf)
	if err1 != nil {
		panic(err1)
	}

	fmt.Println()

	return b64.StdEncoding.EncodeToString(buf.Bytes())
}
