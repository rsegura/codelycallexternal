package cli

import (
	"fmt"
	"strings"
	"github.com/spf13/cobra"
	"github.com/rsegura/codelycallexternal/internal/pokemon"

)
type CobraFn func(cmd *cobra.Command, args []string)




const urlFlag = "url"
const csvFlag = "csv"
func InitPokeCmd(handler pokemon.PokemonHandler) *cobra.Command{
	pokeCmd := &cobra.Command{
		Use: "PokeApi",
		Short: "Print data about pokemons",
		Run: runPokeFn(handler),
	}
	pokeCmd.Flags().StringP(urlFlag, "u", "", "url")
	pokeCmd.Flags().StringP(csvFlag, "c", "", "csv")

	return pokeCmd
}


func runPokeFn(handler pokemon.PokemonHandler) CobraFn {
	return func(cmd *cobra.Command, args []string){
		url,_ := cmd.Flags().GetString(urlFlag)
		name,_ := cmd.Flags().GetString(csvFlag)
		if url == "" || !strings.Contains(url, "pokeapi.co") {
			url = "https://pokeapi.co/api/v2/pokemon?limit=10"
		}
		if name == ""{
			name = "result"
		}
		pokemons := handler.Get(url, name)
		fmt.Println(pokemons.Results)
		
	}
}

