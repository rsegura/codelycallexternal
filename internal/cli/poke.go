package cli

import (
	"os"
	"fmt"
	"log"
	"strings"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"encoding/csv"
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
const csvFlag = "csv"
func InitPokeCmd() *cobra.Command{
	pokeCmd := &cobra.Command{
		Use: "PokeApi",
		Short: "Print data about pokemons",
		Run: runPokeFn(),
	}
	pokeCmd.Flags().StringP(urlFlag, "u", "", "url")
	pokeCmd.Flags().StringP(csvFlag, "c", "", "csv")

	return pokeCmd
}


func runPokeFn() CobraFn {
	return func(cmd *cobra.Command, args []string){
		url,_ := cmd.Flags().GetString(urlFlag)
		csvName,_ := cmd.Flags().GetString(csvFlag)
		if url == "" || !strings.Contains(url, "pokeapi.co") {
			url = "https://pokeapi.co/api/v2/pokemon?limit=10"
		}
		if csvName == ""{
			csvName = "result.csv"
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
		writeCsv(data, csvName)
	}
}

func writeCsv(data pokemonRequest, csvName string){
	file, err := os.Create(csvName)
	defer file.Close()
	checkError("Cannot create file", err)
	writer := csv.NewWriter(file)
	defer writer.Flush()
	checkError("Cannot write to file", err)
	for _, value := range data.Results {
		var record []string
		record = append(record, value.Name)
		record = append(record, value.Url)
		err := writer.Write(record)
		checkError("Cannot write to file", err)
	}
}

func checkError(message string, err error) {
    if err != nil {
        log.Fatal(message, err)
    }
}
