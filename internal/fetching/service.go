package fetching

import(
	"net/http"
	"io/ioutil"
	"encoding/json"
	//"github.com/pkg/errors"
	"github.com/rsegura/codelycallexternal/internal/errors"
)

type PokemonRequest struct{
	Count int `json:"count"`
	Next string `json:"next"`
	Previous string `json:"previous"`	
	Results []struct{
		Name string `json:"name"`
		Url string `json:"url"`
	} `json:"results"`
}


type Service interface{
	FetchPokemons(url string)(PokemonRequest, error)
}

type service struct{

}

func NewService() Service{
	return &service{}
}

func (s *service) FetchPokemons(url string)(PokemonRequest, error){
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
	return data, nil
}

