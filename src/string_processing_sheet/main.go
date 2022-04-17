package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

func (c *Cards) Import() {
	fmt.Println("File name:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	f, err := os.OpenFile(input, os.O_RDONLY, 0444)
	defer f.Close()
	if err != nil {
		fmt.Println("File not found.")
		return
	}
	i := 0
	r := bufio.NewScanner(f)
	for r.Scan() {
		card := r.Text()
		pair := strings.Split(card, "\\s")
		if len(pair) < 2 {
			continue
		}
		i++
		c.data[pair[0]] = pair[1]
		c.definitionSet[pair[1]] = true
	}
	fmt.Printf("%v cards have been loaded\n", i)
}

func (c *Cards) Export() {
	fmt.Println("File name:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	f, err := os.OpenFile(input, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	defer f.Close()
	if err != nil {
		fmt.Printf("error opening file %v \n", input)
		return
	}

	f.WriteString("\n")
	for k, v := range c.data {
		f.WriteString(fmt.Sprintf("%v\\s%v\n", k, v))
	}
	fmt.Printf("%v cards have been saved.\n", len(c.data))
}

func (c *Cards) Ask() {
	fmt.Println("How many times to ask?")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	limit, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}

	// dump keys into slice
	keys := make([]string, 0, len(c.data))
	for k := range c.data {
		keys = append(keys, k)
	}

	i := 0
	idx := 0
	for i < limit {

		if idx == len(keys) {
			idx = 0
		}

		k := keys[idx]
		v := c.data[k]
		fmt.Printf("Print the definition of \"%v\":\n", k)
		input, _ = reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == v {
			fmt.Println("Correct!")
		} else if t, ok := c.definitionSet[input]; ok {
			fmt.Printf("Wrong. The right answer is \"%v\", but your definition is correct for \"%v\".\n", v, t)
		} else {
			fmt.Printf("Wrong. The right answer is \"%v\".\n", v)
		}

		idx++
		i++
	}

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
		case "import":
			cards.Import()
		case "export":
			cards.Export()
		case "ask":
			cards.Ask()
		case "exit":
			fmt.Println("Bye bye!")
			return
		}
	}

}
