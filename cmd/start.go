package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/ko07ga/kadou/lib/redis"
	"github.com/ko07ga/kadou/lib/sqlite3"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "稼働を開始します",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		rc := redis.NewClient(ctx)
		sc := sqlite3.NewClient(ctx)
		defer sc.Close()

		var stampleAt int
		var kadouType string
		sc.Read("SELECT stamped_at,kadou_type from kadous", &stampleAt, &kadouType)
		sc.Exec(fmt.Sprintf("INSERT INTO kadous(stamped_at,kadou_type) VALUES(%d,'開始')", time.Now().UnixMicro()))
		todayStart := fmt.Sprintf(
			"%s-start",
			time.Now().Format("2006-01-02"))
		started, err := rc.Get(todayStart)
		if err != nil {
			panic(err)
		}
		if started != "" {
			fmt.Println("稼働がすでに開始されています")
			return
		}
		fmt.Println("started", started)
		err = rc.Set(todayStart, time.Now().Format("15:04:05"))
		if err != nil {
			panic(err)
		}
		fmt.Println("稼働を開始しました")
	},
}
