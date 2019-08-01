package json


import(
	"encoding/json"
	"github.com/rsegura/codelycallexternal/internal/fetching"
	"github.com/rsegura/codelycallexternal/internal/errors"
	"io/ioutil"


)

type JsonRepo interface{
	SavePokemons(pokemons fetching.PokemonRequest, name string)(error)
}
type repository struct{

}

func NewRepository() JsonRepo{
	return &repository{}
}

func(r *repository) SavePokemons(data fetching.PokemonRequest, name string)(error){
	file, err :=json.MarshalIndent(data, "","")
	if err != nil{
		return errors.WrapDataUnreacheable(err, "error marshaling Json")
	}
	err = ioutil.WriteFile(name + ".json", file, 0644)
	if err != nil{
		return errors.WrapDataUnreacheable(err, "error writing to file %s", name)
	}
	return nil
	/*file, err := os.Create(csvName)
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
	return nil*/
}
