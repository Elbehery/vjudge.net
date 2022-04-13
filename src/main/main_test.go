package main

import (
	"fmt"
	"testing"
)

func TestGreeting(t *testing.T) {
	t.Run("greeting", func(t *testing.T) {
		fmt.Println("hello judge")
	})
}
