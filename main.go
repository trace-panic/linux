package main

import (
	"log"

	"github.com/trace-panic/linux/initializer"
)

func BaseLoop() {
	for {

	}
}

func main() {
	if err := initializer.Init(); err != nil {
		log.Panic("Error configuring OS: ", err)
		return
	}

	BaseLoop()
}
