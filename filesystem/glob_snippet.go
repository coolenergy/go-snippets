package filesystem

import (
	"fmt"
	"path/filepath"
)

func GlobSnippet() {
	goFiles, err := filepath.Glob("./**/*.go")

	if err != nil { // ErrBadPattern
		panic(err)
	}

	for i := 0; i < len(goFiles); i++ {
		fmt.Println(goFiles[i])
	}
}
