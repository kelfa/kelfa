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
	rootCmd.AddCommand(sessionsCmd)
	sessionsCmd.MarkFlagRequired("from")
	sessionsCmd.MarkFlagRequired("to")
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
	a, err := analytics.New(&ds, fromTime, toTime)
	if err != nil {
		return err
	}
	ss := a.GetSessions()
	for _, s := range ss.Sessions {
		if len(s.DataPoints) > 1 {
			spew.Dump(s)
		}
	}
	return nil
}
