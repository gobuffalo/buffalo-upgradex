package cmd

import (
	"fmt"
	"os"

	"github.com/gobuffalo/buffalo-upgradex/cmd/upgradex"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var opts = upgradex.Options{}

// upgradexCmd represents the upgradex command
var upgradexCmd = &cobra.Command{
	Use:     "upgradex",
	Aliases: []string{"install"},
	Short:   "updates Buffalo and/or Pop/Soda as well as your app",
	// DisableFlagParsing: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		pwd, _ := os.Getwd()
		opts.Path = pwd
		err := upgradex.All(opts)
		if errors.Cause(err) == upgradex.ErrNotBuffaloApp {
			fmt.Printf("[WARN] %s", err.Error())
			return nil
		}
		return err
	},
}

func init() {
	upgradexCmd.Flags().BoolVar(&opts.WithSQLite, "sqlite", false, "adds sqlite support")
	upgradexCmd.Flags().BoolVar(&opts.SkipBuffalo, "skip-buffalo", false, "skips updating buffalo")
	upgradexCmd.Flags().BoolVar(&opts.SkipPop, "skip-pop", false, "skips updating pop/soda")
	upgradexCmd.Flags().BoolVar(&opts.SkipApp, "skip-app", false, "skips updating app")
	rootCmd.AddCommand(upgradexCmd)
}
