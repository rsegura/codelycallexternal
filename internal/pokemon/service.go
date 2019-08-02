package pokemon


import(
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/rsegura/codelycallexternal/internal/errors"
)

type Service interface{
	FetchPokemons(url string, name string)(PokemonRequest, error)
}

type service struct{
	repo PokemonRepository
}

func NewPokemonService(repo PokemonRepository) Service{
	return &service{
		repo,
	}
}

func (s *service) FetchPokemons(url string, name string)(PokemonRequest, error){
	res, err := http.Get(url)
	binaryResponse, err := ioutil.ReadAll(res.Body)
	var data PokemonRequest
	var jsonErr error
	if err != nil{
		return PokemonRequest{}, errors.WrapDataUnreacheable(err, "error getting response to %s", url)
	}
	
	jsonErr = json.Unmarshal(binaryResponse, &data)
	if jsonErr != nil {
		return PokemonRequest{}, errors.WrapDataUnreacheable(jsonErr, "error unmarhalling response from %s", url)
	}
	err = s.repo.SavePokemons(data, name)
	if err != nil{
		return PokemonRequest{}, errors.WrapDataUnreacheable(jsonErr, "error saving object from %s", url)
	}
	return data, nil
}