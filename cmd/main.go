package main


import(
	"github.com/spf13/cobra"
	"github.com/rsegura/codelycallexternal/internal/cli"
	"github.com/rsegura/codelycallexternal/internal/pokemon"
	//"github.com/rsegura/codelycallexternal/internal/database/csv"
	"github.com/rsegura/codelycallexternal/internal/database/json"
)

func main(){

	//var pokemonRepository = csv.NewCsvPokemonRepository()
	var pokemonRepository = json.NewJsonPokemonRepository()
	var pokemonService = pokemon.NewPokemonService(pokemonRepository)

	rootCmd := &cobra.Command{Use:"Poke-cli"}
	rootCmd.AddCommand(cli.InitPokeCmd(pokemonService))
	rootCmd.Execute()
}