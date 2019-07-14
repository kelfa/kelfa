package main

import (
	"fmt"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"go.kelfa.io/kelfa/pkg/analytics"
	"go.kelfa.io/kelfa/pkg/dal"
	"go.kelfa.io/kelfa/pkg/dal/objects"
)

func init() {
	rootCmd.AddCommand(cacheCmd)
	cacheCmd.MarkFlagRequired("from")
	cacheCmd.MarkFlagRequired("to")
}

var cacheCmd = &cobra.Command{
	Use:   "cache",
	Short: "see hits rates",
	Args: func(cmd *cobra.Command, args []string) error {
		var err error
		fromTime, err = time.Parse(time.RFC3339, viper.GetString("period_from"))
		if err != nil {
			return fmt.Errorf("%v is not a valid date in the YYYY-MM-DDTHH:mm:SS+ZZ:ZZ format", viper.GetString("period_from"))
		}
		toTime, err = time.Parse(time.RFC3339, viper.GetString("period_to"))
		if err != nil {
			return fmt.Errorf("%v is not a valid date in the YYYY-MM-DDTHH:mm:SS+ZZ:ZZ format", viper.GetString("period_to"))
		}
		return nil
	},
	RunE: cache,
}

func cache(cmd *cobra.Command, args []string) error {
	ds, err := dal.New("filesystem", objects.BackendOptions{Path: viper.GetString("data_folder"), From: fromTime, To: toTime})
	if err != nil {
		return err
	}
	a, err := analytics.New(&ds, fromTime, toTime)
	if err != nil {
		return err
	}
	cs := a.CacheStats()
	spew.Dump(cs)
	return nil
}
