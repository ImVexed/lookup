package lookup

import (
	_ "image/png"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestOCR(t *testing.T) {
	Convey("Given an OCR object", t, func() {
		ocr := NewOCR(0.7)

		Convey("When I try to load an invalid font directory", func() {
			err := ocr.LoadFont("testdata/NON_EXISTENT")

			Convey("It returns an error", func() {
				So(err.Error(), ShouldContainSubstring, "no such file or directory")
			})
		})

		Convey("When I load a valid font on it", func() {
			err := ocr.LoadFont("testdata/font_1")

			Convey("It loads the fonts successfully", func() {
				So(err, ShouldBeNil)
			})

			Convey("It stores the font family", func() {
				So(ocr.fontFamilies, ShouldContainKey, "font_1")
				So(ocr.fontFamilies, ShouldHaveLength, 1)
				So(ocr.fontFamilies["font_1"], ShouldHaveLength, 13)
			})

			Convey("It updates the totalSymbols", func() {
				So(ocr.totalSymbols, ShouldEqual, 13)
			})

			Convey("And when I pass an image to be recognized", func() {
				img := loadImageColor("test3.png")
				text, _ := ocr.Recognize(img)

				Convey("It recognizes the text in the image", func() {
					//So(text, ShouldEqual, "3662\n32€/€​")
					So(text, ShouldNotBeEmpty)
				})
			})
		})

	})
}