package main


import(
	"github.com/spf13/cobra"
	"github.com/rsegura/codelycallexternal/internal/cli"
	"github.com/rsegura/codelycallexternal/internal/fetching"
	"github.com/rsegura/codelycallexternal/internal/storage/csv"
)

func main(){

	repo := csv.NewRepository()
	fetchingService := fetching.NewService()
	rootCmd := &cobra.Command{Use: "Poke-cli"}
	rootCmd.AddCommand(cli.InitPokeCmd(fetchingService, repo))
	rootCmd.Execute()
}