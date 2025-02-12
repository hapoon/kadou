package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "kadouのバージョンを表示します",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("kadou v0.1.0 -- HEAD")
	},
}
