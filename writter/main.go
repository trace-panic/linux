package writter

import "fmt"

type Color struct {
	Reset   string
	Green   string
	Red     string
	Yellow  string
	Blue    string
	Magenta string
	Cyan    string
	White   string
}

var Colors = Color{
	Reset:   "\033[0m",
	Red:     "\033[31m",
	Green:   "\033[32m",
	Yellow:  "\033[33m",
	Blue:    "\033[34m",
	Magenta: "\033[35m",
	Cyan:    "\033[36m",
	White:   "\033[37m",
}

func RedText(message string) string {
	return Colors.Red + message + Colors.Reset
}

func GreenText(message string) string {
	return Colors.Green + message + Colors.Reset
}

func YellowText(message string) string {
	return Colors.Yellow + message + Colors.Reset
}

func BlueText(message string) string {
	return Colors.Blue + message + Colors.Reset
}

func MagentaText(message string) string {
	return Colors.Magenta + message + Colors.Reset
}

func CyanText(message string) string {
	return Colors.Cyan + message + Colors.Reset
}

func WhiteText(message string) string {
	return Colors.White + message + Colors.Reset
}

func ClearConsole() {
	fmt.Print("\033[H\033[2J")
}
