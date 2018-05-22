package lighting

import "image/color"

// Linear represents a linear chain of lights.
//
// The start of a Linear chain will ALWAYS be at Inner, that is, address 0
// when Linear is used is ALWAYS towards Inner.
type Linear struct {
	InnerFern *Fern // Linear's inner fern
	OuterFern *Fern // Linear's outer fern(s)

	// Mapping of LEDs on the chain. This is cleared on every Run().
	LEDs []*color.RGBA
}

// Fern represents a fern.
type Fern struct {
	InnerLinear *Linear   // A Fern's inner linear
	OuterLinear []*Linear // A fern's outer linear(s)
	Arms        [8][5]*color.RGBA
}

// TreeTop represents the lights on the top of the tree.
type TreeTop struct{}

// TreeBase represents the lights at the base of the tree.
type TreeBase struct{}
