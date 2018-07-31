package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Card struct {
	number int
	name   string
	suit   string
}

type Deck []Card

func New() Deck {
	var newDeck Deck
	numbers := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	names := [...]string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}
	suits := [...]string{"Diamond", "Club", "Heart", "Spade"}
	for i, n := range numbers {
		for _, s := range suits {
			newDeck = append(newDeck, Card{number: n, name: names[i], suit: s})
		}
	}
	return newDeck
}

func NewMultiple(n int) Deck {
	var newDeck Deck
	for i := 0; i < n; i++ {
		addition := New()
		newDeck = append(newDeck, addition...)
	}
	return newDeck
}

func (d *Deck) Print() {
	for _, card := range *d {
		if card.suit == "Joker" {
			fmt.Println("Joker")
		} else {
			fmt.Printf("%v of %vs\n", card.name, card.suit)
		}
	}
}

func (d *Deck) Sort(comparator func(i, j int) bool) {
	sort.Slice(*d, comparator)
}

func (d *Deck) DefaultComparator(a int, b int) bool {
	if (*d)[a].number < (*d)[b].number {
		return true
	} else if (*d)[a].number > (*d)[b].number {
		return false
	} else if (*d)[a].suit <= (*d)[b].suit {
		return true
	}
	return false
}

func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(*d), func(i, j int) {
		tempNumber, tempName, tempSuit := (*d)[i].number, (*d)[i].name, (*d)[i].suit
		(*d)[i].number, (*d)[i].name, (*d)[i].suit = (*d)[j].number, (*d)[j].name, (*d)[j].suit
		(*d)[j].number, (*d)[j].name, (*d)[j].suit = tempNumber, tempName, tempSuit
	})
}

func (d *Deck) AddJokers(j int) {
	for i := 0; i < j; i++ {
		*d = append(*d, Card{number: 14, name: "Joker", suit: "Joker"})
	}
}

func (d *Deck) Filter(cards ...string) {
	for i := 0; i < len(cards); i += 2 {
		name := cards[i]
		suit := cards[i+1]
		for j := 0; j < len(*d); j++ {
			if (*d)[j].name == name && (*d)[j].suit == suit {
				*d = append((*d)[:j], (*d)[j+1:]...)
				j--
			}
		}
	}
}

func main() {
	deck := NewMultiple(1)
	deck.Sort(deck.DefaultComparator)
	deck.AddJokers(5)
	deck.Print()
	deck.Filter("9", "Diamond", "10", "Diamond", "King", "Spade")
	deck.Print()
}
