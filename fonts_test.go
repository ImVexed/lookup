package lookup

import (
	_ "image/png"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFontSymbol(t *testing.T) {
	Convey("Given an image and a template to look for", t, func() {
		img := loadImageGray("font_1/0.png")
		fs := newFontSymbol("0", img)
		So(fs.Width, ShouldEqual, img.Bounds().Max.X)
		So(fs.Height, ShouldEqual, img.Bounds().Max.Y)
	})
}

func TestLoadFont(t *testing.T) {
	Convey("Given a font directory", t, func() {
		Convey("When loading the symbols", func() {
			fonts, _ := loadFont("testdata/font_1")

			Convey("It loads all font files", func() {
				So(len(fonts), ShouldEqual, 13)
			})

			Convey("It loads all symbol names correctly", func() {
				var expectedNames = []string{"/", "€", "€", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
				var actualNames []string
				for _, f := range fonts {
					actualNames = append(actualNames, f.Symbol)
				}

				So(actualNames, ShouldResemble, expectedNames)
			})
		})
	})
}
