package puppy

import "fmt"

func Bark() {
	fmt.Println(GetBark())
}

func GetBark() string {
	return "Woof!"
}

func Barks() {
	fmt.Println(GetBarks())
}

func GetBarks() string {
	return "Woof! Woof! Woof!"
}
