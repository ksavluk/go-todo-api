package main

import (
	"log"
	"os"
)

func main() {
	if err := newApp().run(os.Args); err != nil {
		log.Fatalf("[ERROR] %v", err)
	}
}
