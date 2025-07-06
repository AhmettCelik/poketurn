package models

type Ability struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	EffectEntries []struct {
		Effect   string `json:"effect"`
		Language struct {
			Name string `json:"name"`
		}
		ShortEffect string `json:"short_effect"`
	} `json:"effect_entries"`
	FlavorTextEntries []struct {
		FlavorText string `json:"flavor_text"`
		Language   struct {
			Name string `json:"name"`
		}
	} `json:"flavor_text_entries"`
	Pokemons []struct {
		IsHidden bool `json:"is_hidden"`
		Pokemon  struct {
			Name string `json:"name"`
		}
		Slot int `json:"slot"`
	} `json:"pokemon"`
}
