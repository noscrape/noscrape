package noscrape

import (
	"seehuhn.de/go/postscript/funit"
	"seehuhn.de/go/sfnt"
)

func blueValues(font sfnt.Font, r rune, bbox []funit.Int16) []funit.Int16 {
	cmapTable, err := font.CMapTable.GetBest()

	if err != nil || cmapTable == nil {
		panic("cmap could not be found")
	}

	id := cmapTable.Lookup(r)
	ext := font.GlyphBBox(id)

	bottom := ext.LLy
	if bottom < bbox[0] {
		bbox[0] = bottom
	} else if bottom > bbox[1] {
		bbox[1] = bottom
	}

	top := ext.URy
	if top < bbox[2] {
		bbox[2] = top
	} else if top > bbox[3] {
		bbox[3] = top
	}

	return bbox
}
