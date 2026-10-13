package main

import (
	"slices"
	"testing"
)

type Item struct {
	Nome string
	Qtd  int
}

func TestSlice(t *testing.T) {
	got := []Item{{"a", 1}, {"b", 2}}
	esp := []Item{{"a", 1}, {"b", 2}}
	if !slices.Equal(got, esp) {
		t.Errorf("got %v, esperado %v", got, esp)
	}
	if !slices.Equal([]Item{}, nil) {
		t.Error("vazio e nil deveriam ser iguais")
	}
}
