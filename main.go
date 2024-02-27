package main

import (
	"fmt"

	"github.com/CoopHive/coophive/cmd/gdcn"
)

func main() {
	gdcn.Execute()
}

func init() {
	fmt.Printf("GDCN\n")
}
