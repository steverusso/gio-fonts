package freeroad

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
		register(text.Font{}, freeroadTTF)
		register(text.Font{Style: text.Italic}, freeroadItalicTTF)
		register(text.Font{Weight: text.Light}, freeroadLightTTF)
		register(text.Font{Weight: text.Light, Style: text.Italic}, freeroadLightItalicTTF)
		register(text.Font{Weight: text.Bold}, freeroadBoldTTF)
		register(text.Font{Weight: text.Bold, Style: text.Italic}, freeroadBoldItalicTTF)
		register(text.Font{Weight: text.Black}, freeroadBlackTTF)
		register(text.Font{Weight: text.Black, Style: text.Italic}, freeroadBlackItalicTTF)
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
	fnt.Typeface = "Freeroad"
	collection = append(collection, text.FontFace{Font: fnt, Face: face})
}

//go:embed Freeroad-Regular.ttf
var freeroadTTF []byte

//go:embed Freeroad-Italic.ttf
var freeroadItalicTTF []byte

//go:embed Freeroad-Light.ttf
var freeroadLightTTF []byte

//go:embed Freeroad-LightItalic.ttf
var freeroadLightItalicTTF []byte

//go:embed Freeroad-Bold.ttf
var freeroadBoldTTF []byte

//go:embed Freeroad-BoldItalic.ttf
var freeroadBoldItalicTTF []byte

//go:embed Freeroad-Black.ttf
var freeroadBlackTTF []byte

//go:embed Freeroad-BlackItalic.ttf
var freeroadBlackItalicTTF []byte
