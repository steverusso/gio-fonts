package vegur

import (
	_ "embed"
	"fmt"
	"sync"

	"gioui.org/font/opentype"
	"gioui.org/text"
)

var (
	once       sync.Once
	collection []text.FontFace
)

func Collection() []text.FontFace {
	once.Do(func() {
		register(text.Font{}, vegurOTF)
		register(text.Font{Weight: text.Bold}, vegurBoldOTF)
		register(text.Font{Weight: text.Light}, vegurLightOTF)
		// Ensure that any outside appends will not reuse the backing store.
		n := len(collection)
		collection = collection[:n:n]
	})
	return collection
}

func register(fnt text.Font, data []byte) {
	face, err := opentype.Parse(data)
	if err != nil {
		panic(fmt.Errorf("failed to parse font: %v", err))
	}
	fnt.Typeface = "Vegur"
	collection = append(collection, text.FontFace{Font: fnt, Face: face})
}

//go:embed Vegur-Regular.otf
var vegurOTF []byte

//go:embed Vegur-Bold.otf
var vegurBoldOTF []byte

//go:embed Vegur-Light.otf
var vegurLightOTF []byte
