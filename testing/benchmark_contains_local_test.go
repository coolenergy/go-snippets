package testing

import (
	"bytes"
	"strings"
	"testing"
)

func BenchmarkLocalContainsString(b *testing.B) {
	helloWorldStr := "Hello World"
	helloStr := "Hello"
	b.ResetTimer()
	b.ReportAllocs()
	var r bool
	for i := 0; i < b.N; i++ {
		r = strings.Contains(helloWorldStr, helloStr)
	}

	result = r
}

func BenchmarkLocalContainsBytes(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	helloWorldBytes := []byte("Hello World")
	var r bool
	for i := 0; i < b.N; i++ {
		r = bytes.Contains(helloWorldBytes, []byte("Hello"))
	}
	result = r
}

func BenchmarkLocalContainsBytes2(b *testing.B) {
	// This is a useful benchmark for HTTP clients
	// a lot of times we retrieve responses using http package, which returns response as []byte
	// is it worth converting to str?
	// Conclusion: it is not worth converting to string if all the APIs you need can also handle
	// []byte, for example if you are gonna search, you can use bytes.Contains, no need for conversion
	// if you only are gonna use strings.Contains()

	b.ResetTimer()
	b.ReportAllocs()
	helloWorldBytes := []byte("Hello World")
	var r bool
	for i := 0; i < b.N; i++ {
		r = strings.Contains(string(helloWorldBytes), "Hello")
	}
	result = r
}

func BenchmarkLocalContainsBytes3(b *testing.B) {
	// This is a useful benchmark for HTTP clients
	// a lot of times we retrieve responses using http package, which returns response as []byte
	// is it worth converting to str?
	// Conclusion: it is not worth converting to string if all the APIs you need can also handle
	// []byte, for example if you are gonna search, you can use bytes.Contains, no need for conversion
	// if you only are gonna use strings.Contains()

	b.ResetTimer()
	b.ReportAllocs()

	var r bool
	for i := 0; i < b.N; i++ {
		r = strings.Contains(string(helloWorldBytes), helloStr)
	}
	result = r
}
