package testing

import (
	"bytes"
	"go4.org/strutil"
	"io/ioutil"
	"strings"
	"testing"
)

// https://medium.com/@felipedutratine/in-golang-should-i-work-with-bytes-or-strings-8bd1f5a7fd48
// As stated in that article:
// Strings are faster for searches (contains, index, compare) purpose.

var searchBytes = []byte("https://www.w3.org/2000/svg")
var searchString = "https://www.w3.org/2000/svg"

func stringContains(message string, search string) bool {

	searchIndex := 0

	for i := 0; i < len(message); i++ {
		if i+len(search)-searchIndex > len(message) {
			return false
		}

		if message[i] == search[searchIndex] {
			searchIndex++

			if searchIndex == len(search) {
				return true
			}
		}
	}

	return false
}
func byteSliceContains(message []byte, search []byte) bool {

	searchIndex := 0

	for i := 0; i < len(message); i++ {
		if i+len(search)-searchIndex > len(message) {
			return false
		}

		if message[i] == search[searchIndex] {
			searchIndex++

			if searchIndex == len(search) {
				return true
			}
		}
	}

	return false
}

var contentBytes, _ = ioutil.ReadFile("test_files/big_file.html")
var contentString = string(contentBytes)

func BenchmarkByteSliceContains(b *testing.B) {
	for i := 0; i < b.N; i++ {
		byteSliceContains(contentBytes, searchBytes)
	}
}

func BenchmarkStringContains(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringContains(contentString, "https://www.w3.org/2000/svg")
	}
}

func BenchmarkStringContainsFramework(b *testing.B) {
	contentString := string(contentBytes)
	for i := 0; i < b.N; i++ {
		strings.Contains(contentString, "https://www.w3.org/2000/svg")
	}
}

func BenchmarkByteArrayContainsFramework(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bytes.Contains(contentBytes, searchBytes)
	}
}

func BenchmarkGo4Org(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strutil.ContainsFold(contentString, searchString)
	}
}
