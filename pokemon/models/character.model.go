package models

import "encoding/json"

// CharacterPage contient la liste de Pokémon renvoyée par Tyradex.
// L'API renvoie directement un tableau JSON, sans objet "results" autour.
type CharacterPage struct {
	Results []Character
}

// UnmarshalJSON adapte le tableau de l'API au champ Results utilisé par le service.
func (page *CharacterPage) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &page.Results)
}

// Character représente les champs principaux d'un Pokémon dans Tyradex.
type Character struct {
	ID         int      `json:"pokedex_id"`
	Generation int      `json:"generation"`
	Category   string   `json:"category"`
	Name       Names    `json:"name"`
	Sprites    Sprites  `json:"sprites"`
	Types      []Type   `json:"types"`
	Stats      Stats    `json:"stats"`
	Talents    []Talent `json:"talents"`
	Height     string   `json:"height"`
	Weight     string   `json:"weight"`
}

// Names contient les noms dans les langues proposées par l'API.
type Names struct {
	French   string `json:"fr"`
}

// Sprites contient les URL des illustrations du Pokémon.
type Sprites struct {
	Regular string             `json:"regular"`
}

// GigantamaxSprites contient les illustrations Gigamax, quand elles existent.
type GigantamaxSprites struct {
	Regular string `json:"regular"`
}

// Type est un type élémentaire et son icône.
type Type struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}

// Stats contient les statistiques de base.
type Stats struct {
	HP             int `json:"hp"`
	Attack         int `json:"atk"`
	Defense        int `json:"def"`
	SpecialAttack  int `json:"spe_atk"`
	SpecialDefense int `json:"spe_def"`
	Speed          int `json:"vit"`
}

// Talent décrit une capacité et précise si elle est cachée (talent caché).
type Talent struct {
	Name   string `json:"name"`
	Hidden bool   `json:"tc"`
}
