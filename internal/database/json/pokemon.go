package json

import(
	"encoding/json"
	"github.com/rsegura/codelycallexternal/internal/pokemon"
	"github.com/rsegura/codelycallexternal/internal/errors"
	"io/ioutil"
)

type pokemonRepository struct{

}


func NewJsonPokemonRepository() pokemon.PokemonRepository {
	return &pokemonRepository{
	}
}

func (p *pokemonRepository) SavePokemons(data pokemon.PokemonRequest, name string)(error){
	file, err :=json.MarshalIndent(data, "","")
	if err != nil{
		return errors.WrapDataUnreacheable(err, "error marshaling Json")
	}
	err = ioutil.WriteFile(name + ".json", file, 0644)
	if err != nil{
		return errors.WrapDataUnreacheable(err, "error writing to file %s", name)
	}
	return nil
}