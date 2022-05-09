package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use: "loko",
}

func init() {
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(startCmd)
}

func Execute() {
	rootCmd.Execute()
}
