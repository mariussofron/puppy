package puppy

import (
	"fmt"

	bigdog "github.com/mariussofron/big-dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof! Woof!"
}

func BigBark() string {
	return bigdog.BigBarks(Bark())
}

func BigBarks() string {
	return bigdog.BigBarks(Barks())
}

func FromV1() {
	fmt.Println("This is version v1.0.0")
}