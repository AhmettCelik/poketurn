package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/AhmettCelik/poketurn/internal/models"
)

func FetchPokemon(name string, id int) (models.Pokemon, error) {
	url := "https://pokeapi.co/api/v2/pokemon/"
	if id == 0 {
		url = url + name
	} else {
		url = url + strconv.Itoa(id)
	}

	pkmn := models.Pokemon{}

	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error getting response : %v\n", err)
		return pkmn, err
	}

	pkmn, err = decodeResponse[models.Pokemon](res.Body)
	if err != nil {
		return pkmn, err
	}

	if res.StatusCode > 299 {
		return pkmn, fmt.Errorf("Response failed with status code: %d\n", res.StatusCode)
	}

	return pkmn, nil
}
