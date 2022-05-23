package scooter

import (
	"fmt"
	"testing"
)

func TestRandom(t *testing.T) {
	text, err := Random()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(text)
}

func TestRandomFull(t *testing.T) {
	response, err := RandomFull()
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v\n", response)
}

func TestDaily(t *testing.T) {
	text, err := Daily()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(text)
}

func TestDailyFull(t *testing.T) {
	response, err := DailyFull()
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v\n", response)
}

func TestQuoteById(t *testing.T) {
	text, err := QuoteById(1)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(text)
}

func TestFullQuoteById(t *testing.T) {
	response, err := QuoteByIdFull(1)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v\n", response)
}
