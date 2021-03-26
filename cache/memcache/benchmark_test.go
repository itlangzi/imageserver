package memcache

import (
	"testing"

	"github.com/itlangzi/imageserver"
	cachetest "github.com/itlangzi/imageserver/cache/_test"
	"github.com/itlangzi/imageserver/testdata"
)

func BenchmarkGet(b *testing.B) {
	for _, tc := range []struct {
		name string
		im   *imageserver.Image
	}{
		{"Small", testdata.Small},
		{"Medium", testdata.Medium},
		{"Large", testdata.Large},
	} {
		b.Run(tc.name, func(b *testing.B) {
			cch := newTestCache(b)
			cachetest.BenchmarkGet(b, cch, 1, tc.im) // memcached is unstable with more parallelism
		})
	}
}
