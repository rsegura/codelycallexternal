package main


import(
	"github.com/spf13/cobra"
	"github.com/rsegura/codelycallexternal/internal/cli"
)

func main(){
	rootCmd := &cobra.Command{Use: "Poke-cli"}
	rootCmd.AddCommand(cli.InitPokeCmd())
	rootCmd.Execute()
}