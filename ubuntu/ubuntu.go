package ubuntu

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
		register(text.Font{}, ubuntuTTF)
		register(text.Font{Style: text.Italic}, ubuntuItalicTTF)
		register(text.Font{Weight: text.Bold}, ubuntuBoldTTF)
		register(text.Font{Weight: text.Medium}, ubuntuMediumTTF)
		register(text.Font{Variant: "Mono"}, ubuntuMonoTTF)
		register(text.Font{Variant: "Mono", Weight: text.Bold}, ubuntuMonoBoldTTF)
		register(text.Font{Variant: "Mono", Weight: text.Bold, Style: text.Italic}, ubuntuMonoBoldItalicTTF)
		register(text.Font{Variant: "Mono", Style: text.Italic}, ubuntuMonoItalicTTF)
		// Ensure that any outside appends will not reuse the backing store.
		n := len(collection)
		collection = collection[:n:n]
	})
	return collection
}

func register(fnt text.Font, ttf []byte) {
	face, err := opentype.Parse(ttf)
	if err != nil {
		panic(fmt.Errorf("failed to parse font: %v", err))
	}
	fnt.Typeface = "Ubuntu"
	collection = append(collection, text.FontFace{Font: fnt, Face: face})
}

//go:embed Ubuntu-R.ttf
var ubuntuTTF []byte

//go:embed Ubuntu-RI.ttf
var ubuntuItalicTTF []byte

//go:embed Ubuntu-B.ttf
var ubuntuBoldTTF []byte

//go:embed Ubuntu-M.ttf
var ubuntuMediumTTF []byte

//go:embed UbuntuMono-R.ttf
var ubuntuMonoTTF []byte

//go:embed UbuntuMono-B.ttf
var ubuntuMonoBoldTTF []byte

//go:embed UbuntuMono-Italic.ttf
var ubuntuMonoItalicTTF []byte

//go:embed UbuntuMono-BoldItalic.ttf
var ubuntuMonoBoldItalicTTF []byte
