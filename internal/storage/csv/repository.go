package csv

import(
	"os"
	"encoding/csv"
	"log"
	"github.com/rsegura/codelycallexternal/internal/fetching"


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
		checkError("Cannot create file", err)
		return err
	}
	
	writer := csv.NewWriter(file)
	defer writer.Flush()
	for _, value := range data.Results {
		var record []string
		record = append(record, value.Name)
		record = append(record, value.Url)
		err := writer.Write(record)
		if err != nil{
			checkError("Cannot write to file", err)
			return err
		}
	}
	return nil
}
func checkError(message string, err error) {
    if err != nil {
        log.Fatal(message, err)
    }
}