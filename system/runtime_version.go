package system

import (
	"fmt"
	"runtime"
)

func runtimeVersion() {
	fmt.Printf("Go version: %s", runtime.Version())
}
