package asap

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
		register(text.Font{}, asapOTF)
		register(text.Font{Weight: text.Bold}, asapBoldOTF)
		register(text.Font{Style: text.Italic}, asapItalicOTF)
		register(text.Font{Weight: text.Bold, Style: text.Italic}, asapBoldItalicOTF)
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
	fnt.Typeface = "Asap"
	collection = append(collection, text.FontFace{Font: fnt, Face: face})
}

//go:embed Asap-Regular.otf
var asapOTF []byte

//go:embed Asap-Bold.otf
var asapBoldOTF []byte

//go:embed Asap-Italic.otf
var asapItalicOTF []byte

//go:embed Asap-BoldItalic.otf
var asapBoldItalicOTF []byte
