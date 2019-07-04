package cli

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/spf13/cobra"
)
type CobraFn func(cmd *cobra.Command, args []string)


type pokemonRequest struct{
	Count int `json:"count"`
	Next int `json:"next"`
	Previous int `json:"previous"`	
	Results []struct{
		Name string `json:"name"`
		Url string `json:"url"`
	} `json:"results"`
}

const urlFlag = "url"
func InitPokeCmd() *cobra.Command{
	pokeCmd := &cobra.Command{
		Use: "PokeApi",
		Short: "Print data about pokemons",
		Run: runPokeFn(),
	}
	pokeCmd.Flags().StringP(urlFlag, "u", "", "url")

	return pokeCmd
}


func runPokeFn() CobraFn {
	return func(cmd *cobra.Command, args []string){
		url,_ := cmd.Flags().GetString(urlFlag)
		if url == "" {
			fmt.Println("entramos")
			url = "https://pokeapi.co/api/v2/pokemon?limit=10"
		}
		var data pokemonRequest
		var jsonErr error
		res, err := http.Get(url)
		if err != nil{
			fmt.Println(err)
		}
		if jsonErr != nil {
			fmt.Println(jsonErr)
		}
		binaryResponse, err := ioutil.ReadAll(res.Body)
		jsonErr = json.Unmarshal(binaryResponse, &data)
		fmt.Println(data)
	}
}
