package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/fxkk/goscooter/pkg/scooter"
	"github.com/spf13/cobra"
)

// dailyCmd represents the daily command
var dailyCmd = &cobra.Command{
	Use:   "daily",
	Short: "Daily Scooter Quote",
	Run: func(cmd *cobra.Command, args []string) {
		if Full != true {
			text, err := scooter.Daily()
			if err != nil {
				fmt.Println("Failed to get random quote")
			}
			fmt.Println(text)
		} else {
			response, err := scooter.DailyFull()
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
	rootCmd.AddCommand(dailyCmd)
}
