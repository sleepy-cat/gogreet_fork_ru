package greet

import "fmt"
import "github.com/sleepy-cat/gogreet/format"

func Greet(name string) string {
	return fmt.Sprintf(format.GreetingFormat(), name)
}
