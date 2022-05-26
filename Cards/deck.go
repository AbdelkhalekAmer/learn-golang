package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}

	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Queen", "Prince", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}

	return cards
}

func newDeckFromFile(filename string) deck {
	data, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error: " + err.Error())
		os.Exit(1)
	}

	cardsFromFile := strings.Split(string(data), ",")

	return deck(cardsFromFile)
}

func (dck deck) print() {
	for _, card := range dck {
		fmt.Println(card)
	}
}

func deal(dck deck, handSize int) (deck, deck) {
	return dck[:handSize], dck[handSize:]
}

func (dck deck) toString() string {
	return strings.Join(dck, ",")
}

func (dck deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(dck.toString()), 0666)
}

func (dck deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())

	r := rand.New(source)

	for i := range dck {
		randomIndex := r.Intn(len(dck) - 1)
		dck[i], dck[randomIndex] = dck[randomIndex], dck[i]
	}
}
