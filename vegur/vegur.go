package vegur

import (
	_ "embed"
	"fmt"
	"sync"

	"gioui.org/font/opentype"
	"gioui.org/text"
	"github.com/steverusso/gio-fonts/vegur/vegurbold"
	"github.com/steverusso/gio-fonts/vegur/vegurlight"
	"github.com/steverusso/gio-fonts/vegur/vegurregular"
)

var (
	once       sync.Once
	collection []text.FontFace
)

func Collection() []text.FontFace {
	once.Do(func() {
		register(text.Font{}, vegurregular.OTF)
		register(text.Font{Weight: text.Bold}, vegurbold.OTF)
		register(text.Font{Weight: text.Light}, vegurlight.OTF)
		// Ensure that any outside appends will not reuse the backing store.
		n := len(collection)
		collection = collection[:n:n]
	})
	return collection
}

func register(fnt text.Font, data []byte) {
	face, err := opentype.Parse(data)
	if err != nil {
		panic(fmt.Errorf("failed to parse vegur font %v: %v", fnt, err))
	}
	fnt.Typeface = "Vegur"
	collection = append(collection, text.FontFace{Font: fnt, Face: face})
}
