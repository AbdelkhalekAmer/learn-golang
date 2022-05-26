package main

func main() {
	cards := newDeckFromFile("demo_cards.txt")

	cards.shuffle()

	hand, _ := deal(cards, 5)

	hand.print()
}
