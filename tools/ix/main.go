package main

import (
	"encoding/json"
	"os"

	app "ix/internal"
)

func main() {
	if err := app.Run(os.Args[1:], os.Stdout); err != nil {
		_ = json.NewEncoder(os.Stderr).Encode(map[string]string{"error": err.Error()})
		os.Exit(1)
	}
}
