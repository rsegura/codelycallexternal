package pokemon



type PokemonHandler interface {
	Get(url string, name string) (PokemonRequest)
}


type pokemonHandler struct {
	pokemonService PokemonService
}

func NewPokemonHandler(pokemonService PokemonService) PokemonHandler {
	return &pokemonHandler{
		pokemonService,
	}
}


func(h *pokemonHandler) Get(url, name string) (PokemonRequest) {
	pokemons, _ := h.pokemonService.FetchPokemons(url, name)
	return pokemons
}