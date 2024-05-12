package cmd

import (
	"github.com/RodgerCodes/Go-zig-and-python-playground/sorter/pkg"
	"github.com/spf13/cobra"
)

var sortDownloadsCmd = &cobra.Command{
	Use:   "downloads",
	Short: "Sort files in downloads directory/folder",
	Long:  "Sort files in downloads directory/folder",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := pkg.SortDownloads()
		if err != nil {
			return err
		}

		return nil
	},
}
