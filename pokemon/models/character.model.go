package models

// CharacterPage : une page de la liste des personnages.
// Seul le tableau "results" est utile ici (l'objet "info" sert à la pagination).
type CharacterPage struct {
	Results []Character `json:"results"`
}

// Character : fiche d'un personnage renvoyée par Tyradex - Une API Pokémon en français.
// Seuls les champs utiles au site sont déclarés.
type Character struct {
	ID              int      `json:"pokedex_id"`
	Category        string   `json:"category"`
	Name            string   `json:"name"`
	Language        string   `json:"fr"`
	Graphic_element string   `json:"sprites"`
	Type_graphic    string   `json:"regular"`
	Types           []string `json:"types"`
	Stats           []string `json:"stats"`
	Skills          []string `json:"talents"`
}

// Place : un lieu (origine ou lieu actuel du personnage).
type Place struct {
	Name string `json:"name"`
}
