package main

import (
	"fmt"
	"os"
	"path/filepath"

	plugin "github.com/bots-garden/wasm-tinygo-pdk"
)

func main() {
	if filepath.Ext(os.Args[0]) == ".wasm" {
		// CLI mode
		value, err := Handle([]byte(os.Args[1]))
		fmt.Println(string(value), err)
	} else {
		// Plugin mode
		plugin.SetHandle(Handle)
	}
}

func Handle(param []byte) ([]byte, error) {

	return []byte("Hello " + string(param)), nil

}
