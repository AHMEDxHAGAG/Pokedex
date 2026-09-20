package pokeapi

var Location struct {
	PokemonEncounters []pokemonEncounters `json:"pokemon_encounters"`
}

type pokemonEncounters struct {
	Pokemon namedAPIResource `json:"pokemon"`
}

type Pokemon struct {
	Name           string  `json:"name"`
	Height         int     `json:"height"`
	Weight         int     `json:"weight"`
	BaseExperience int     `json:"base_experience"`
	Stats          []stats `json:"stats"`
	Types          []types `json:"types"`
}

type stats struct {
	Stat     namedAPIResource `json:"stat"`
	BaseStat int              `json:"base_stat"`
}

type types struct {
	Type namedAPIResource `json:"type"`
}

type namedAPIResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type result struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

var LocationArea struct {
	Next    string   `json:"next"`
	Prev    string   `json:"previous"`
	Results []result `json:"results"`
}
