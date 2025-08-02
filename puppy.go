package puppy

import (
	"github.com/brunorgavloski/dog"
)

func Bark() string {
	return "woof!"
}

func Bark() string {
	return dog.WhenGrowUp(Bark())
}
func BigBark() string {
	return dog.WhenGrowUp(Bark())
}

func BigBark() string {
	return dog.WhenGrowUp(Bark())
}
