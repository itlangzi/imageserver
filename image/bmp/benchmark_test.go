package bmp

import (
	"testing"

	"github.com/itlangzi/imageserver"
	imageserver_image_test "github.com/itlangzi/imageserver/image/_test"
	_ "github.com/itlangzi/imageserver/image/jpeg"
	"github.com/itlangzi/imageserver/testdata"
)

func Benchmark(b *testing.B) {
	enc := &Encoder{}
	params := imageserver.Params{}
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
			imageserver_image_test.BenchmarkEncoder(b, enc, tc.im, params)
		})
	}
}
