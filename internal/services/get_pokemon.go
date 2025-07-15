package services

import (
	"github.com/AhmettCelik/poketurn/internal/api"
	"github.com/AhmettCelik/poketurn/internal/models"
)

func GetPokemonByName(name string) (models.Pokemon, error) {
	return api.FetchPokemon(name, 0)
}

func GetPokemonByID(id int) (models.Pokemon, error) {
	return api.FetchPokemon("", id)
}

func GetAllPokemons() (models.AllPokemons, error) {
	return api.FetchAllPokemon()
}
