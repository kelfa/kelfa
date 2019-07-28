package main

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"go.kelfa.io/kelfa/pkg/analytics"
	"go.kelfa.io/kelfa/pkg/dal"
	"go.kelfa.io/kelfa/pkg/dal/objects"
)

func init() {
	rootCmd.AddCommand(sessionsCmd)
}

var sessionsCmd = &cobra.Command{
	Use:   "sessions",
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
	RunE: sessions,
}

func sessions(cmd *cobra.Command, args []string) error {
	ds, err := dal.New("filesystem", objects.BackendOptions{Path: viper.GetString("data_folder"), From: fromTime, To: toTime})
	if err != nil {
		return err
	}
	mode, err := analytics.ParseMode(viper.GetString("mode"))
	if err != nil {
		return err
	}
	a, err := analytics.New(&ds, fromTime, toTime, mode, viper.GetDuration("session_inactivity_timeout"))
	if err != nil {
		return err
	}
	var sessionsAmount int
	var cs int
	for _, ds := range a.Data {
		for _, s := range ds.Sessions {
			if s.Crawler {
				cs++
			}
			sessionsAmount++
		}
	}
	fmt.Printf("a total of %v sessions have been registered\n", sessionsAmount)
	fmt.Printf("of those, %v sessions are made by crawlers\n", cs)
	return nil
}
