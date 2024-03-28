package testing

import (
	"bytes"
	"strings"
	"testing"
)

var helloWorldStr = "Hello World"
var helloStr = "Hello"
var helloWorldBytes = []byte("Hello World")
var helloBytes = []byte("Hello")
var result interface{}

func BenchmarkContainsString(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	var r bool
	for i := 0; i < b.N; i++ {
		r = strings.Contains(helloWorldStr, helloStr)
	}

	result = r
}

func BenchmarkContainsBytes(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	var r bool
	for i := 0; i < b.N; i++ {
		r = bytes.Contains(helloWorldBytes, helloWorldBytes)
	}
	result = r
}
