package test

import (
	"testing"

	"github.com/AhmettCelik/poketurn/internal/services"
)

func TestGetPokemonByName_Success(t *testing.T) {
	pkmnName := "bulbasaur"
	got, err := services.GetPokemonByName(pkmnName)
	if err != nil {
		t.Fatal(err)
	}

	if got.Name != pkmnName {
		t.Fatalf("Name = %s, want %s", got.Name, pkmnName)
	}
}
