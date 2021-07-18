package main

import (
	"fmt"
	"os"

	"github.com/filipeandrade6/vigia-go/internal/gravacao/core"
)

func main() {
	if err := core.Main(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}
