package cmd

import (
	"context"

	"github.com/ko07ga/kadou/lib/sqlite3"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "初期化を実行します",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		sc := sqlite3.NewClient(ctx)
		defer sc.Close()
		stmt := `
		CREATE TABLE kadous(
			stamped_at int not null primary key,
			kadou_type string)`
		sc.Exec(stmt)
	},
}
