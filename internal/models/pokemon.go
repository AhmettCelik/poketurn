package models

type Pokemon struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	BaseEXP   int    `json:"base_experience"`
	Height    int    `json:"height"`
	Weight    int    `json:"weight"`
	Abilities []struct {
		Ability struct {
			Name string `json:"name"`
		} `json:"ability"`
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
	} `json:"abilities"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
		}
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}
