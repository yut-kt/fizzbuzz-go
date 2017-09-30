package main

import (
	"testing"
	"strings"
)

func TestFizz(t *testing.T) {
	if !strings.EqualFold("FizzBuzz", FizzBuzz(15)) {
		t.Errorf("Correct! Wow, I hope require FizzBuzz")
	}
	if !strings.EqualFold("Fizz", FizzBuzz(3)) {
		t.Errorf("Correct! Wow, I hope require Fizz")
	}
	if !strings.EqualFold("Buzz", FizzBuzz(5)) {
		t.Errorf("Correct! Wow, I hope require Buzz")
	}
	if !strings.EqualFold("98", FizzBuzz(98)) {
		t.Errorf("Correct! Wow, I hope require 98")
	}
}
