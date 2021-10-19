package main

import (
	"fmt"
)

func main() {
	//declares a slice
	myCourses := make([]string, 5, 10)
	myCourses2 := []string{"Docker", "Kubernetes", "Python", "5", "1", "2", "5", "1", "1", "2", "5", "1", "1", "2", "5", "1"}
	//array capacity doubles when it is expanded.
	for _, i:= range myCourses2 {
		myCourses = append(myCourses, i)
		fmt.Println("\nLength is: %d. \nCapacity is: %d", len(myCourses), cap(myCourses))
	}
	myCourses = append(myCourses, myCourses2...)
	fmt.Println(myCourses)
	

	
}
