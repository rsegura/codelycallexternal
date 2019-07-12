package csv

import(
	"os"
	"encoding/csv"
	"github.com/rsegura/codelycallexternal/internal/fetching"
	"github.com/rsegura/codelycallexternal/internal/errors"



)

type CsvRepo interface{
	SavePokemons(pokemons fetching.PokemonRequest, csvName string)(error)
}
type repository struct{

}

func NewRepository() CsvRepo{
	return &repository{}
}

func(r *repository) SavePokemons(data fetching.PokemonRequest, csvName string)(error){
	file, err := os.Create(csvName)
	defer file.Close()
	if err != nil{
		return errors.WrapDataUnreacheable(err, "error creating file %s", csvName)
	}
	
	writer := csv.NewWriter(file)
	defer writer.Flush()
	for _, value := range data.Results {
		var record []string
		record = append(record, value.Name)
		record = append(record, value.Url)
		err := writer.Write(record)
		if err != nil{
			return errors.WrapDataUnreacheable(err, "error writing to file %s", csvName)
		}
	}
	return nil
}
