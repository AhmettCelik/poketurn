package services

import (
	"github.com/AhmettCelik/poketurn/internal/api"
	"github.com/AhmettCelik/poketurn/internal/models"
)

func GetPokemonByName(name string) (models.Pokemon, error) {
	pkmn, err := api.FetchPokemon(name, 0)
	if err != nil {
		return pkmn, err
	}

	return pkmn, nil
}
