package main

import (
	"fmt"
	"log"

	"github.com/trace-panic/linux/commands"
	"github.com/trace-panic/linux/console"
	"github.com/trace-panic/linux/initializer"
)

func BaseLoop() {
	for {
		fmt.Print("patrick@dell:~$ ")
		input := console.ReadInput()

		commands.RunCommand(input)
	}
}

func main() {
	if err := initializer.Init(); err != nil {
		log.Panic("Error configuring OS: ", err)
		return
	}

	BaseLoop()
}
