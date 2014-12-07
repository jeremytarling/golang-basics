package main


import "fmt"
import "sort"


// declare empty slices (for sorting the cards and properties)		
var cardOrder[]int

						
func main() {	
	
// append each card in the map to the cardOrder slice
	for card := range majorArcana {
	    cardOrder = append(cardOrder, card)
	}
	sort.Ints(cardOrder)
	
// iterate through the map of maps
	for cardNum := range cardOrder {
	    fmt.Println("\nKey", cardNum)
		// todo - sort keys and values for individual card properties
		for key, value := range majorArcana[cardNum] {
			fmt.Println(key + ":", value)
		}	
	}
		
}


// map of tarot major arcana properties */
var majorArcana = map[int]map[string]string {
		0 : map[string]string{
		"name":"The Fool",
		"hebrew letter":"Aleph",
		"hebrew letter value":"1",
		"hebrew letter meaning":"Ox",
		"astrological association":"Air",
	},
		1 : map[string]string{
		"name":"The Magician",
		"hebrew letter":"Beth",
		"hebrew letter value":"2",
		"hebrew letter meaning":"House",
		"astrological association":"Mercury",
	},
		2 : map[string]string{
		"name":"The High Priestess",
		"hebrew letter":"Gimmel",
		"hebrew letter value":"2",
		"hebrew letter meaning":"Camel",
		"astrological association":"Moon",
	},
}