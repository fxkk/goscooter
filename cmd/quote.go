package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/fxkk/goscooter/pkg/scooter"
	"github.com/spf13/cobra"
)

func validateArguments(args []string) error {
	if len(args) == 1 {
		if _, err := strconv.Atoi(args[0]); err != nil {
			return errors.New("Quote id must be an integer")
		}
	} else {
		return errors.New("Provide exactly one quote id")
	}

	return nil
}

// quoteCmd represents the quote command
var quoteCmd = &cobra.Command{
	Use:   "quote",
	Short: "Get Quote by id",
	Run: func(cmd *cobra.Command, args []string) {
		err := validateArguments(args)
		if err != nil {
			fmt.Println(err)
			return
		}
		id, _ := strconv.Atoi(args[0])
		if Full != true {
			text, err := scooter.QuoteById(id)
			if err != nil {
				fmt.Printf("Failed to get quote: %d\n", id)
			}
			fmt.Println(text)
		} else {
			response, err := scooter.QuoteByIdFull(id)
			if err != nil {
				fmt.Printf("Failed to get quote: %d\n", id)
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
	rootCmd.AddCommand(quoteCmd)
}
