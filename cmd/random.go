package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/fxkk/goscooter/pkg/scooter"
	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Random Scooter Quote",
	Run: func(cmd *cobra.Command, args []string) {
		if Full != true {
			text, err := scooter.Random()
			if err != nil {
				fmt.Println("Failed to get random quote")
			}
			fmt.Println(text)
		} else {
			response, err := scooter.RandomFull()
			if err != nil {
				fmt.Println("Failed to get random quote")
			}

			s, err := json.Marshal(response)
			if err != nil {
				fmt.Println("Could not marshal response to json")
			}
			fmt.Print(string(s))
		}

	},
}

func init() {
	rootCmd.AddCommand(randomCmd)
}
