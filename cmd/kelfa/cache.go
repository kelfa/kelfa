package main

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(cacheCmd)
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

var fromTime time.Time
var toTime time.Time

func cache(cmd *cobra.Command, args []string) error {
	/*	mode, err := analytics.ParseMode(viper.GetString("mode"))
		if err != nil {
			return err
		}
		a, err := analytics.New(&ds, fromTime, toTime, mode, viper.GetDuration("session_inactivity_timeout"))
		if err != nil {
			return err
		}
		cs := a.CacheStats()
		spew.Dump(cs)
	*/
	ds.CacheStats()
	return nil
}
