// implementa rotinas de utiliario
// https://pkg.go.dev/path/filepath#example-Walk
//go:build !windows && !plan9

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Printf("\n" + path)
		return nil
	})

}
