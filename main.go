package main

import (
	"fmt"
	"log"

	"github.com/trace-panic/linux/initializer"
)

func main() {
	if err := initializer.CreateConfig(); err != nil {
		log.Fatal("Error creating configuration: ", err)
	}

	cfg := initializer.GetConfig()
	fmt.Printf("Configuration: %+v\n", cfg)
}
