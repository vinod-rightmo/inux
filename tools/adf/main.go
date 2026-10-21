package main

import (
	"fmt"
	"os"

	app "adf/internal"
)

func main() {
	if err := app.Run(os.Args[1:], os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, "adf:", err)
		os.Exit(1)
	}
}
