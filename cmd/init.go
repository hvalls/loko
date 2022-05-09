package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Execute request",
	Run: func(cmd *cobra.Command, args []string) {

		content := []byte("services: \n")
		err := os.WriteFile("./docker-compose.yml", content, 0644)
		check(err)
	},
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
