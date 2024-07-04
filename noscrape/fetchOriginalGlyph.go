package noscrape

import (
	"fmt"
	"seehuhn.de/go/sfnt"
	"seehuhn.de/go/sfnt/cff"
	"seehuhn.de/go/sfnt/glyf"
)

// fetchOriginalGlyph retrieves the original glyph from the provided font based on the given rune.
// It converts the TrueType glyph outline to a CFF glyph outline and returns the resulting CFF glyph.
func fetchOriginalGlyph(font sfnt.Font, originalRune rune) *cff.Glyph {
	// Get the best cmap table from the font
	originalCmap, err := font.CMapTable.GetBest()
	if err != nil {
		panic("cmap not found")
	}

	// Lookup the glyph ID for the original rune
	originId := originalCmap.Lookup(originalRune)
	// Get the original outlines from the font
	origOutlines := font.Outlines.(*glyf.Outlines)

	// Create a new CFF glyph with the name and width of the original glyph
	cffGlyph := cff.NewGlyph(font.GlyphName(originId), font.GlyphWidth(originId))

	// Get the original glyph from the outlines
	origGlyf := origOutlines.Glyphs[originId]

	var g glyf.SimpleGlyph
	var ok bool
	if origGlyf != nil {
		g, ok = origGlyf.Data.(glyf.SimpleGlyph)
	}

	if !ok {
		return cffGlyph
	}

	// Decode the glyph information
	glyphInfo, err := g.Decode()
	if err != nil {
		panic(fmt.Sprintf("error decoding glyph: %v", err))
	}

	// Convert TrueType contours to CFF contours
	for _, cc := range glyphInfo.Contours {
		var extended glyf.Contour
		var prev glyf.Point
		onCurve := true
		for _, cur := range cc {
			if !onCurve && !cur.OnCurve {
				extended = append(extended, glyf.Point{
					X:       (cur.X + prev.X) / 2,
					Y:       (cur.Y + prev.Y) / 2,
					OnCurve: true,
				})
			}
			extended = append(extended, cur)
			prev = cur
			onCurve = cur.OnCurve
		}
		n := len(extended)

		var offs int
		for i := 0; i < len(extended); i++ {
			if extended[i].OnCurve {
				offs = i
				break
			}
		}

		// Move to the starting point of the contour
		cffGlyph.MoveTo(float64(extended[offs].X), float64(extended[offs].Y))

		i := 0
		for i < n {
			i0 := (i + offs) % n
			if !extended[i0].OnCurve {
				panic("not on curve")
			}
			i1 := (i0 + 1) % n
			if extended[i1].OnCurve {
				if i == n-1 {
					break
				}
				// Draw a line to the next on-curve point
				cffGlyph.LineTo(float64(extended[i1].X), float64(extended[i1].Y))
				i++
			} else {
				// Convert TrueType curves to CFF curves
				i2 := (i1 + 1) % n
				cffGlyph.CurveTo(
					float64(extended[i0].X)/3+float64(extended[i1].X)*2/3,
					float64(extended[i0].Y)/3+float64(extended[i1].Y)*2/3,
					float64(extended[i1].X)*2/3+float64(extended[i2].X)/3,
					float64(extended[i1].Y)*2/3+float64(extended[i2].Y)/3,
					float64(extended[i2].X),
					float64(extended[i2].Y))
				i += 2
			}
		}
	}

	return cffGlyph
}
