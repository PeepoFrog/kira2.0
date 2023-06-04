package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mrlutik/kira2.0/internal/types"
)

const version = "v0.0.1"

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "version":
			fmt.Println(version)
		default:
			fmt.Printf("Unknown command: %s\n", os.Args[1])
		}
	} else {
		fmt.Println("No command provided.")
	}
	t := types.Test{Text: "Hello from kira2_launcher"}
	log.Printf("Entering internal: %v", t.Text)
}
