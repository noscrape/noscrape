package noscrape

import (
	"seehuhn.de/go/sfnt"
	"seehuhn.de/go/sfnt/cff"
	"seehuhn.de/go/sfnt/glyf"
)

func fetchOriginalGlyph(font sfnt.Font, originalRune rune) *cff.Glyph {
	originalCmap, err := font.CMapTable.GetBest()

	if err != nil {
		panic("cmap not found")
	}

	originId := originalCmap.Lookup(originalRune)
	origOutlines := font.Outlines.(*glyf.Outlines)

	cffGlyph := cff.NewGlyph(font.GlyphName(originId), font.GlyphWidth(originId))

	origGlyf := origOutlines.Glyphs[originId]

	var g glyf.SimpleGlyph
	var ok bool
	if origGlyf != nil {
		g, ok = origGlyf.Data.(glyf.SimpleGlyph)
	}

	if !ok {
		return cffGlyph
	}

	glyphInfo, err := g.Decode()

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
				cffGlyph.LineTo(float64(extended[i1].X), float64(extended[i1].Y))
				i++
			} else {
				// See the following link for converting truetype outlines
				// to CFF outlines:
				// https://pomax.github.io/bezierinfo/#reordering
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
