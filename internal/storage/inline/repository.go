package inline

import(
	"github.com/rsegura/codelycallexternal/internal/fetching"
)

type InlineRepo interface{
	SavePokemons(pokemons fetching.PokemonRequest, csvName string)(error)
}
type repository struct{

}

func NewRepository() InlineRepo{
	return &repository{}
}


func(r *repository) SavePokemons(data fetching.PokemonRequest, name string) (error){

	for _, value := range data.Results {
		println(value.Name)
	
	}
	return nil
}
