package cmd

import (
	"errors"
	"fmt"
	"loko/service"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "add",
	Short: "Add service",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("service name is required")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		definition, err := service.GetDefinition(name)
		check(err)

		f, err := os.OpenFile("docker-compose.yml", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		check(err)
		defer f.Close()

		_, err = f.WriteString("\n")
		check(err)

		lines := strings.Split(definition, "\n")
		for _, line := range lines {
			_, err = f.WriteString(fmt.Sprintf("  %s\n", line))
			check(err)
		}

		fmt.Println(fmt.Sprintf("service %s added successfully.", name))
	},
}
