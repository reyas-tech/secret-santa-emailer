package main

import (
	"fmt"
	"math/rand"
	"slices"
)

type santa struct {
	name   string
	giftee string
	nopair string
}

func main() {
	// Create secret santa group
	santas := []santa{
		{name: "user1", nopair: "user3"},
		{name: "user2"},
		{name: "user3", nopair: "user1"},
		{name: "user4"},
		{name: "user5"},
	}
	santas = pairSantas(santas) // assign giftee to each santa
	fmt.Println(santas)         // every santa should have unique giftee
}

// Assign a giftee to each santa
func pairSantas(santas []santa) []santa {
	for { // regenerate pairs until all pairs are valid
		giftees := giftees(santas) // get pool of unpaired giftees
		for i, santa := range santas {
			var giftee int
			for {
				if len(giftees) > 2 {
					giftee = rand.Intn(len(giftees))                                      // randomly select a giftee
					if giftees[giftee] != santa.name && giftees[giftee] != santa.nopair { // if giftee is current santa or nopair of current santa, regenerate
						break
					}
				} else if len(giftees) == 2 {
					giftee = rand.Intn(len(giftees))   // randomly select a giftee
					if giftees[giftee] != santa.name { // if giftee is current santa, regenerate
						break
					}
				} else { // if only one giftee left in pool
					giftee = 0
					break
				}
			}
			santas[i].giftee = giftees[giftee]                 // set giftee for current santa
			giftees = slices.Delete(giftees, giftee, giftee+1) // remove giftee from giftees pool
		}
		if checkPairs(santas) { // if valid santa/giftee pairs
			return santas // return updated santas with giftees
		}
	}
}

// Copy names from santas group into giftees pool
func giftees(santas []santa) []string {
	giftees := make([]string, len(santas))
	for i := range giftees {
		giftees[i] = santas[i].name
	}
	return giftees
}

// Check if each santa has a valid giftee assigned
func checkPairs(santas []santa) bool {
	valid := true
	for _, santa := range santas {
		if santa.name == santa.giftee || santa.giftee == santa.nopair {
			valid = false
			break
		}
	}
	return valid
}
