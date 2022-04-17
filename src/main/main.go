package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Card struct {
	Term       string
	Definition string
}

func (c *Card) Validate(answer string) bool {
	return c.Definition == answer
}

func New(term, definition string) *Card {
	return &Card{
		Term:       term,
		Definition: definition,
	}
}

func main() {

	cards := []*Card{}
	termsSet := map[string]bool{}
	definitionSet := map[string]string{}
	reader := bufio.NewReader(os.Stdin)
	// read the number
	var n int
	fmt.Println("Input the number of cards:")
	fmt.Scanln(&n)

	i := 1
	for i <= n {
		// read the term
		fmt.Printf("The term for card #%v:\n", i)
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		term := strings.TrimSpace(line)
		for termsSet[term] {
			fmt.Printf("The term \"%v\" already exists. Try again:\n", term)
			line, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
			term = strings.TrimSpace(line)
		}
		termsSet[term] = true

		// read the definition
		fmt.Printf("The definition for card #%v:\n", i)
		line, err = reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		definition := strings.TrimSpace(line)
		for definitionSet[definition] != "" {
			fmt.Printf("The definition \"%v\" already exists. Try again:\n", definition)
			line, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
			definition = strings.TrimSpace(line)
		}
		definitionSet[definition] = term

		// create the card and append
		card := New(term, definition)
		cards = append(cards, card)

		i++
	}

	for _, c := range cards {
		fmt.Printf("Print the definition of \"%v\":\n", c.Term)
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		answer := strings.TrimSpace(line)
		if c.Validate(answer) {
			fmt.Print("Correct!\n")
		} else {
			if t, ok := definitionSet[answer]; ok {
				fmt.Printf("Wrong. The right answer is \"%v\", but your definition is correct for \"%v\".\n", c.Definition,t)
			} else {
				fmt.Printf("Wrong. The right answer is \"%v\".\n", c.Definition)
			}
		}
	}
}
