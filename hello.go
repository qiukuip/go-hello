package hello

import "fmt"

func Hello(name string) string {
	if len(name) == 0 {
		name = "world"
	}
	return fmt.Sprintf("hello %s!\n", name)
}
