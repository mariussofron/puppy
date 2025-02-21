package puppy

import (
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