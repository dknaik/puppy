package puppy

import (
	"github.com/dknaik/dogs"
)

func Bark() string {
	return "woof!"
}

func Barks() string {
	return "woof! woof! woof!"
}

func BigBark() string {
	return dogs.WhenGrowsUp(Bark())
}

func BigBarks() string {
	return dogs.WhenGrowsUp(Barks())
}
