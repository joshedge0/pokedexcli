package pokeapi

type PokemonApiBody struct {
	Name           string  `json:"name"`
	BaseExperience float32 `json:"base_experience"`
}
