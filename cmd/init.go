package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize docker-compose.yml",
	Run: func(cmd *cobra.Command, args []string) {
		createNewFile("docker-compose.yml", "services: \n")
	},
}

func createNewFile(fileName, content string) {
	if _, err := os.Stat(fileName); errors.Is(err, os.ErrNotExist) {
		err = os.WriteFile(fileName, []byte(content), 0644)
		check(err)
		fmt.Printf("Created %s file\n", fileName)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
