package page

import (
	"math"
	"math/rand"
	"testing"
)

const (
	maxKeySize    = 100
	maxRecordSize = 1_000
)

var (
	pageSize = int(math.Round(2 * float64(Size) / 3)) // two-third of the max size
)

func generateRecordCell() RecordCell {
	keySize := 1 + rand.Intn(maxKeySize-1)
	recordSize := 1 + rand.Intn(maxRecordSize-1)
	key, record := make([]byte, keySize), make([]byte, recordSize)
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}
	_, err = rand.Read(record)
	if err != nil {
		panic(err)
	}

	return RecordCell{
		Key:    key,
		Record: record,
	}
}

func generatePage() *Page {
	p, err := load(make([]byte, pageSize))
	if err != nil {
		panic(err)
	}
	record := generateRecordCell()
	for err = p.StoreRecordCell(record); err != ErrPageFull; err = p.StoreRecordCell(record) {
		record = generateRecordCell()
	}
	return p
}

func BenchmarkPage_Defragment(b *testing.B) {
	data := generatePage().data
	bytes := make([]byte, len(data))
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		copy(bytes, data)
		page := &Page{ // have to recreate a page with initial data
			data: bytes,
		}
		page.Defragment()
	}
}
