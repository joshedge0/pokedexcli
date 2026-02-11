package pokeapi

type ResLocationDeep struct {
	//EncounterMethodRates string   `json:"encounter_method_rates"`
	//GameIndex            int      `json:"game_index"`
	//Id                   int      `json:"id"`
	Location Location `json:"location"`
	Name     string   `json:"name"`
	//Names                []string `json:"names"`
	PokemonEncounters []struct {
		Pokemon Pokemon `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type ResLocationShallow struct {
	Count    int        `json:"count"`
	Next     *string    `json:"next"`
	Previous *string    `json:"previous"`
	Results  []Location `json:"results"`
}
