package main

import (
	"fmt"

	"github.com/bacalhau-project/generic-dcn/cmd/gdcn"
)

func main() {
	gdcn.Execute()
}

func init() {
	fmt.Printf("GDCN\n")
}
