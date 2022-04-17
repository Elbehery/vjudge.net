package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Cards struct {
	data          map[string]string
	definitionSet map[string]bool
}

func NewCards() *Cards {
	return &Cards{
		data:          make(map[string]string),
		definitionSet: make(map[string]bool),
	}
}

func (c *Cards) Remove() {
	fmt.Println("Which card?")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if _, ok := c.data[input]; ok {
		delete(c.data, input)
		fmt.Println("The card has been removed.")
		return
	}
	fmt.Printf("Can't remove \"%v\": there is no such card.\n", input)
}

func (c *Cards) Add() {
	// capture term from user
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("The card:")
	term, _ := reader.ReadString('\n')
	term = strings.TrimSpace(term)
	for c.data[term] != "" {
		fmt.Printf("The card \"%v\" already exists. Try again:\n", term)
		line, _ := reader.ReadString('\n')
		term = strings.TrimSpace(line)
	}

	// capture definition from user
	fmt.Println("The definition of the card:")
	definition, _ := reader.ReadString('\n')
	definition = strings.TrimSpace(definition)
	for c.definitionSet[definition] {
		fmt.Printf("The definition \"%v\" already exists. Try again:\n", definition)
		line, _ := reader.ReadString('\n')
		definition = strings.TrimSpace(line)
	}

	// print msg to user
	fmt.Println(ConfirmMsg(term, definition))
	// add to indices
	c.data[term] = definition
	c.definitionSet[definition] = true
}

func ConfirmMsg(term, definition string) string {
	return fmt.Sprintf("The pair (\"%v\":\"%v\") has been added.\n", term, definition)
}

func main() {

	// validation data structure
	cards := NewCards()

	for {
		//capture input cmd
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Input the action (add, remove, import, export, ask, exit):")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "add":
			cards.Add()
		case "remove":
			cards.Remove()
		case "exit":
			fmt.Println("Bye bye!")
			return
		}
	}

}
