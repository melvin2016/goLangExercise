package main

import (
	"fmt"       /* fmt package for formatting */
	"io/ioutil" /* io utility package */
	"math/rand" /* rand type from math package to perform random number operations*/
	"os"        /* os functionalities package */
	"strings"   /* strings package for string manipulation */
	"time"      /* Time Package */
)

// Creating new type of deck with base type of slice of string
type deck []string

// Receiver function to print all items in a deck of cards.
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Function to create a new deck of cards.
func newDeck() deck {
	// New Deck
	cards := deck{}

	// Card Suites
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}

	//Card Values
	cardValues := []string{"Ace", "One", "Two", "Three", "Four"}

	// Function to make new deck with suites and values
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	// returning new deck of cards
	return cards
}

// Function to deal cards.
func deal(d deck, num int) (deck, deck) {
	cardsInHand := d[:num]
	cardsLeft := d[num:]
	return cardsInHand, cardsLeft
}

// Reciever function to make a deck to string.
func (d deck) toString() string {
	return strings.Join(d, ",")
}

// Reciever function to write to disk.
func (d deck) writeToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// Function to read from disk.
func readFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	return deck(strings.Split(string(bs), ","))
}

// Function to shuffle the cards
func (d deck) shuffle() {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	for i := range d {
		rn := r.Intn(len(d) - 1)
		d[i], d[rn] = d[rn], d[i]
	}
}
