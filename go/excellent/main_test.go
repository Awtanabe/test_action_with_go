package main

import (
	"log"
	"testing"
)


func TestEventOrOdd(t *testing.T) {
	
	result := EventOrOdd(10)

	log.Print("hoge")
	log.Print("hoge")
	if result != "event" {
		t.Errorf("expected event actual %s", result)
	}
}