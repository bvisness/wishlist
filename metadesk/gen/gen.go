package main

import (
	"io"
	"os"

	"github.com/bvisness/wishlist/metadesk"
)

// Don't run this directly! Run `go generate` on the root.
func main() {
	refFile, err := os.Open("reference.mdesk")
	if err != nil {
		panic(err)
	}

	reference, err := io.ReadAll(refFile)
	if err != nil {
		panic(err)
	}

	os.Remove("md.go") // ignore errors, nothing will go wrong
	f, err := os.Create("md.go")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	_, err = f.Write(metadesk.GenBindings(string(reference)))
	if err != nil {
		panic(err)
	}
}
