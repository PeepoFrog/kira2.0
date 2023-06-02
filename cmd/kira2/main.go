package main

import (
	"log"

	"github.com/mrlutik/kira2.0/internal/types"
)

func main() {

	t := types.Test{Text: "Hello from kira2_launcher"}
	log.Printf("Entering internal: %v", t.Text)
}
