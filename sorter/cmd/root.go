package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "Sorter",
	Short: "Sort your file properly",
	Long:  `Sort your file properly`,
}

func init() {
	rootCmd.AddCommand(sortCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
