package metropolis

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
		register(text.Font{}, metropolisOTF)
		register(text.Font{Style: text.Italic}, metropolisItalicOTF)
		register(text.Font{Weight: text.Thin}, metropolisThinOTF)
		register(text.Font{Weight: text.Thin, Style: text.Italic}, metropolisThinItalicOTF)
		register(text.Font{Weight: text.ExtraLight}, metropolisExtraLightOTF)
		register(text.Font{Weight: text.ExtraLight, Style: text.Italic}, metropolisExtraLightItalicOTF)
		register(text.Font{Weight: text.Light}, metropolisLightOTF)
		register(text.Font{Weight: text.Light, Style: text.Italic}, metropolisLightItalicOTF)
		register(text.Font{Weight: text.Medium}, metropolisMediumOTF)
		register(text.Font{Weight: text.Medium, Style: text.Italic}, metropolisMediumItalicOTF)
		register(text.Font{Weight: text.SemiBold}, metropolisSemiBoldOTF)
		register(text.Font{Weight: text.SemiBold, Style: text.Italic}, metropolisSemiBoldItalicOTF)
		register(text.Font{Weight: text.Bold}, metropolisBoldOTF)
		register(text.Font{Weight: text.Bold, Style: text.Italic}, metropolisBoldItalicOTF)
		register(text.Font{Weight: text.ExtraBold}, metropolisExtraBoldOTF)
		register(text.Font{Weight: text.ExtraBold, Style: text.Italic}, metropolisExtraBoldItalicOTF)
		register(text.Font{Weight: text.Black}, metropolisBlackOTF)
		register(text.Font{Weight: text.Black, Style: text.Italic}, metropolisBlackItalicOTF)
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
	fnt.Typeface = "Metropolis"
	collection = append(collection, text.FontFace{Font: fnt, Face: face})
}

//go:embed Metropolis-Regular.otf
var metropolisOTF []byte

//go:embed Metropolis-RegularItalic.otf
var metropolisItalicOTF []byte

//go:embed Metropolis-Thin.otf
var metropolisThinOTF []byte

//go:embed Metropolis-ThinItalic.otf
var metropolisThinItalicOTF []byte

//go:embed Metropolis-ExtraLight.otf
var metropolisExtraLightOTF []byte

//go:embed Metropolis-ExtraLightItalic.otf
var metropolisExtraLightItalicOTF []byte

//go:embed Metropolis-Light.otf
var metropolisLightOTF []byte

//go:embed Metropolis-LightItalic.otf
var metropolisLightItalicOTF []byte

//go:embed Metropolis-Medium.otf
var metropolisMediumOTF []byte

//go:embed Metropolis-MediumItalic.otf
var metropolisMediumItalicOTF []byte

//go:embed Metropolis-SemiBold.otf
var metropolisSemiBoldOTF []byte

//go:embed Metropolis-SemiBoldItalic.otf
var metropolisSemiBoldItalicOTF []byte

//go:embed Metropolis-Bold.otf
var metropolisBoldOTF []byte

//go:embed Metropolis-BoldItalic.otf
var metropolisBoldItalicOTF []byte

//go:embed Metropolis-ExtraBold.otf
var metropolisExtraBoldOTF []byte

//go:embed Metropolis-ExtraBoldItalic.otf
var metropolisExtraBoldItalicOTF []byte

//go:embed Metropolis-Black.otf
var metropolisBlackOTF []byte

//go:embed Metropolis-BlackItalic.otf
var metropolisBlackItalicOTF []byte
