package go_nanocolors

import "testing"

func BenchmarkColor(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Red("red")
	}
}

func BenchmarkNestedColors(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Underline(Bold(Red("Underline and Bold and Red")))
	}
}

func BenchmarkRgb(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		RGB(10, 30, 255)("Hello world!")
	}
}

var rgbTest = RGB(10, 30, 255)

func BenchmarkPreDefinedRgb(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		rgbTest("Hello world!")
	}
}
