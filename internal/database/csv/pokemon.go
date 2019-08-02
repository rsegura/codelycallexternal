package csv

import(
	"os"
	"encoding/csv"
	"github.com/rsegura/codelycallexternal/internal/pokemon"
	"github.com/rsegura/codelycallexternal/internal/errors"
)

type pokemonRepository struct{

}


func NewCsvPokemonRepository() pokemon.PokemonRepository {
	return &pokemonRepository{
	}
}

func (p *pokemonRepository) SavePokemons(data pokemon.PokemonRequest, name string)(error){
	file, err := os.Create(name+".csv")
	defer file.Close()
	if err != nil{
		return errors.WrapDataUnreacheable(err, "error creating file %s", name)
	}
	
	writer := csv.NewWriter(file)
	defer writer.Flush()
	for _, value := range data.Results {
		var record []string
		record = append(record, value.Name)
		record = append(record, value.Url)
		err := writer.Write(record)
		if err != nil{
			return errors.WrapDataUnreacheable(err, "error writing to file %s", name)
		}
	}
	return nil
}