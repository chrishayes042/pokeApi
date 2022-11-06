package pokemonStruct

type PokemonResponse struct {
	Text string `json:"flavor_text"`
	PokemonTextData []PokemonTextData `json:"flavor_text_entries"`
	Happiness PokemonData `json:"base_happiness"`
	CapRate PokemonData `json:"capture_rate"`
}
type Response struct {
	Name string `json:"name"`
	Url string `json:"url"`
	Pokemon []Pokemon `json:"pokemon_entries"`
}

type PokemonTextData struct {
	Text string `json:"flavor_text"`
}

type PokemonData struct {
	Happiness int `json:"base_happiness"`
	CapRate int `json:"capture_rate"`
}

type Pokemon struct {
	EntryNo int `json:"entry_number"`
	Species PokemonSpecies `json:"pokemon_species"`	
}

type PokemonSpecies struct {
	Name string `json:"name"`
	Url string `json:"url"`

}