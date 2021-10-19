package main

import (
	"fmt"
	"os"
)

func main() {
	name := os.Getenv("USERNAME")
	course := "Docker Deep Dive"

	fmt.Println("\nHi", name, "you're currently catching", course)

	course = changeCourse(course)

	fmt.Println("\nYou are now watching course", course)
}

func changeCourse(course string) string {
	course = "First look"
	fmt.Println("\nTrying to change your course to", course)

	return course
}
