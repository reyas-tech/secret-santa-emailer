package main

import (
	"fmt"
	"math/rand"
	"slices"
)

type santa struct {
	name   string
	giftee string
}

func main() {
	// Create secret santa group
	santas := []santa{
		{name: "user1"},
		{name: "user2"},
		{name: "user3"},
		{name: "user4"},
		{name: "user5"},
	}
	santas = pairSantas(santas) // assign giftee to each santa
	fmt.Println(santas)         // every santa should have unique giftee
}

// Assign a giftee to each santa
func pairSantas(santas []santa) []santa {
	giftees := giftees(santas) // get pool of unpaired giftees
	for i := range santas {
		var giftee int
		for {
			giftee = rand.Intn(len(giftees)) // randomly select a giftee
			if giftee != i {                 // if giftee is current santa, regenerate
				break
			}
		}
		santas[i].giftee = giftees[giftee]                 // set giftee for current santa
		giftees = slices.Delete(giftees, giftee, giftee+1) // remove giftee from giftees pool
	}
	return santas // return updated santas with giftees
}

// Copy names from santas group into giftees pool
func giftees(santas []santa) []string {
	giftees := make([]string, len(santas))
	for i := range giftees {
		giftees[i] = santas[i].name
	}
	return giftees
}
