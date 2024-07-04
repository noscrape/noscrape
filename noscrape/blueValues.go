package noscrape

import (
	"seehuhn.de/go/postscript/funit"
	"seehuhn.de/go/sfnt"
)

// blueValues updates the BlueValues array in the private dictionary of a Type1 font.
// It computes the bottom and top bounds of a given rune's glyph bounding box and updates the BlueValues accordingly.
// It returns the updated BlueValues slice.
func blueValues(font sfnt.Font, r rune, bbox []funit.Int16) []funit.Int16 {
	// Get the best cmap table from the font
	cmapTable, err := font.CMapTable.GetBest()
	if err != nil || cmapTable == nil {
		panic("cmap could not be found")
	}

	// Lookup the glyph ID for the rune
	id := cmapTable.Lookup(r)
	// Get the bounding box for the glyph
	ext := font.GlyphBBox(id)

	// Update the bottom bounds in the BlueValues array
	bottom := ext.LLy
	if bottom < bbox[0] {
		bbox[0] = bottom
	} else if bottom > bbox[1] {
		bbox[1] = bottom
	}

	// Update the top bounds in the BlueValues array
	top := ext.URy
	if top < bbox[2] {
		bbox[2] = top
	} else if top > bbox[3] {
		bbox[3] = top
	}

	// Return the updated BlueValues array
	return bbox
}
