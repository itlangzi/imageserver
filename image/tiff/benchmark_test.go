package tiff

import (
	"testing"

	"github.com/itlangzi/imageserver"
	imageserver_image_test "github.com/itlangzi/imageserver/image/_test"
	_ "github.com/itlangzi/imageserver/image/jpeg"
	"github.com/itlangzi/imageserver/testdata"
)

func Benchmark(b *testing.B) {
	for _, tc := range []struct {
		name string
		im   *imageserver.Image
	}{
		{"Small", testdata.Small},
		{"Medium", testdata.Medium},
		{"Large", testdata.Large},
		{"Huge", testdata.Huge},
	} {
		b.Run(tc.name, func(b *testing.B) {
			imageserver_image_test.BenchmarkEncoder(b, &Encoder{}, tc.im, imageserver.Params{})
		})
	}
}
