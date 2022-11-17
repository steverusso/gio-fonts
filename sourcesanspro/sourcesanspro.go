package sourcesanspro

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
		register(text.Font{}, regularTTF)
		register(text.Font{Style: text.Italic}, italicTTF)
		register(text.Font{Weight: text.ExtraLight}, extraLightTTF)
		register(text.Font{Weight: text.ExtraLight, Style: text.Italic}, extraLightItalicTTF)
		register(text.Font{Weight: text.Light}, lightTTF)
		register(text.Font{Weight: text.Light, Style: text.Italic}, lightItalicTTF)
		register(text.Font{Weight: text.SemiBold}, semiBoldTTF)
		register(text.Font{Weight: text.SemiBold, Style: text.Italic}, semiBoldItalicTTF)
		register(text.Font{Weight: text.Bold}, boldTTF)
		register(text.Font{Weight: text.Bold, Style: text.Italic}, boldItalicTTF)
		register(text.Font{Weight: text.Black}, blackTTF)
		register(text.Font{Weight: text.Black, Style: text.Italic}, blackItalicTTF)
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
	fnt.Typeface = "SourceSansPro"
	collection = append(collection, text.FontFace{Font: fnt, Face: face})
}

var (
	//go:embed SourceSansPro-Regular.ttf
	regularTTF []byte
	//go:embed SourceSansPro-Italic.ttf
	italicTTF []byte
	//go:embed SourceSansPro-ExtraLight.ttf
	extraLightTTF []byte
	//go:embed SourceSansPro-ExtraLightItalic.ttf
	extraLightItalicTTF []byte
	//go:embed SourceSansPro-Light.ttf
	lightTTF []byte
	//go:embed SourceSansPro-LightItalic.ttf
	lightItalicTTF []byte
	//go:embed SourceSansPro-Semibold.ttf
	semiBoldTTF []byte
	//go:embed SourceSansPro-SemiboldItalic.ttf
	semiBoldItalicTTF []byte
	//go:embed SourceSansPro-Bold.ttf
	boldTTF []byte
	//go:embed SourceSansPro-BoldItalic.ttf
	boldItalicTTF []byte
	//go:embed SourceSansPro-Black.ttf
	blackTTF []byte
	//go:embed SourceSansPro-BlackItalic.ttf
	blackItalicTTF []byte
)
