package main

import "fmt"

type Card struct {
	number int
	name string
	hebrewLetter string
	hebrewMeaning string
	hebrewNumber int
	correspondence string
}


func main() {
	
	fool := Card {
		number:0,
		name:"The Fool",
		hebrewLetter:"Aleph",
		hebrewMeaning:"Ox",
		hebrewNumber:1,
		correspondence:"Air",
	}
	
	fmt.Println("key", fool.number) 
	fmt.Println("name:", fool.name) 
	fmt.Println("Hebrew letter:", fool.hebrewLetter) 
	fmt.Println("Hebrew letter meaning:", fool.hebrewMeaning) 
	fmt.Println("Hebrew number:", fool.hebrewNumber) 
	fmt.Println("correspondence:", fool.correspondence) 
	
}