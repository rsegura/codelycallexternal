package cli

import (
	"fmt"
	"strings"
	"github.com/spf13/cobra"
	"github.com/rsegura/codelycallexternal/internal/fetching"
	"github.com/rsegura/codelycallexternal/internal/storage/csv"

)
type CobraFn func(cmd *cobra.Command, args []string)




const urlFlag = "url"
const csvFlag = "csv"
func InitPokeCmd(service fetching.Service, repository csv.CsvRepo) *cobra.Command{
	pokeCmd := &cobra.Command{
		Use: "PokeApi",
		Short: "Print data about pokemons",
		Run: runPokeFn(service, repository),
	}
	pokeCmd.Flags().StringP(urlFlag, "u", "", "url")
	pokeCmd.Flags().StringP(csvFlag, "c", "", "csv")

	return pokeCmd
}


func runPokeFn(service fetching.Service, repository csv.CsvRepo) CobraFn {
	return func(cmd *cobra.Command, args []string){
		url,_ := cmd.Flags().GetString(urlFlag)
		csvName,_ := cmd.Flags().GetString(csvFlag)
		if url == "" || !strings.Contains(url, "pokeapi.co") {
			url = "https://pokeapi.co/api/v2/pokemon?limit=10"
		}
		if csvName == ""{
			csvName = "result.csv"
		}
		pokemons, err := service.FetchPokemons(url)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(pokemons)
		repository.SavePokemons(pokemons, csvName)
	}
}

