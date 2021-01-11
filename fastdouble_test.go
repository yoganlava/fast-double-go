package fastdouble

import (
	"strconv"
	"testing"
	"unsafe"
)

func BenchmarkFastDouble(b *testing.B) {
	var res float64
	test := []byte("1.7976931348623157e+308")
	for i := 0; i < b.N; i++ {
		ParseNumber(uintptr(unsafe.Pointer(&test[0])), &res)
	}
}

func BenchmarkParseFloat64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strconv.ParseFloat("1.7976931348623157e+308", 64)
	}
}
