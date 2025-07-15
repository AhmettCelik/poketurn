package models

type AllPokemons struct {
	Results []struct {
		Name string `json:"name"`
	} `json:"results"`
}
