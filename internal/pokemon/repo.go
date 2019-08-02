package pokemon

type PokemonRepository interface{
	SavePokemons(pokemons PokemonRequest, name string)(error)
}