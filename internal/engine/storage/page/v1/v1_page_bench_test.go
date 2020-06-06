package v1

import (
	"testing"

	"github.com/tomarrell/lbadd/internal/engine/storage/page"
)

var result interface{}

func Benchmark_Page_StoreCell(b *testing.B) {
	_p, _ := load(make([]byte, PageSize))
	_ = _p.StoreCell(page.Cell{
		Key:    []byte{0xAA},
		Record: []byte{0xCA, 0xFE, 0xBA, 0xBE},
	})
	_ = _p.StoreCell(page.Cell{
		Key:    []byte{0xFF},
		Record: []byte{0xCA, 0xFE, 0xBA, 0xBE},
	})
	_ = _p.StoreCell(page.Cell{
		Key:    []byte{0x11},
		Record: []byte{0xCA, 0xFE, 0xBA, 0xBE},
	})

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		data := make([]byte, PageSize)
		copy(data, _p.data)
		p, _ := load(data)
		b.StartTimer()

		_ = p.StoreCell(page.Cell{
			Key:    []byte{0xDD},
			Record: []byte{0xCA, 0xFE, 0xBA, 0xBE},
		})
	}
}

func Benchmark_Page_Load(b *testing.B) {
	_p, _ := load(make([]byte, PageSize))
	_ = _p.StoreCell(page.Cell{
		Key:    []byte{0xAA},
		Record: []byte{0xCA, 0xFE, 0xBA, 0xBE},
	})
	_ = _p.StoreCell(page.Cell{
		Key:    []byte{0xFF},
		Record: []byte{0xCA, 0xFE, 0xBA, 0xBE},
	})
	_ = _p.StoreCell(page.Cell{
		Key:    []byte{0x11},
		Record: []byte{0xCA, 0xFE, 0xBA, 0xBE},
	})

	var r page.Page

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		p, _ := Load(_p.data)

		r = p
	}

	result = r
}
