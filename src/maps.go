package main

import (
	"fmt"
)

func main() {

	leagueTitles := make(map[string]int);
	leagueTitles["Sunderland"] = 6
	leagueTitles["NewCastle"] = 4


	recentHead2Head := map[string]int {
		"Sunderland": 5,
		"Newcastle": 0,
	}

	recentHead2Head["things"] = 5

	for key, value := range recentHead2Head {
		fmt.Printf("\nKey is %v Value is: %v", key, value)
	}

	delete(recentHead2Head, "things")

	fmt.Printf("\n Recent Head 2 Head: %v", recentHead2Head)
	fmt.Printf("\nLeague titles: %v\nRecent head to heads: %v\n", leagueTitles, recentHead2Head)
}