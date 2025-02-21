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

func BigBark(s string) string {
	return bigdog.BigBarks(Bark())
}

func BigBarks(s string) string {
	return bigdog.BigBarks(Barks())
}