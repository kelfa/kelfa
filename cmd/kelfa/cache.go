package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"

	"go.kelfa.io/kelfa/pkg/dal"
	"go.kelfa.io/kelfa/pkg/dal/objects"
)

func init() {
	rootCmd.AddCommand(cacheCmd)
	cacheCmd.Flags().StringVarP(&from, "from", "f", "", "from")
	cacheCmd.Flags().StringVarP(&to, "to", "t", "", "to")
}

var from string
var to string
var fromTime time.Time
var toTime time.Time

var cacheCmd = &cobra.Command{
	Use:   "cache",
	Short: "see hits rates",
	Args: func(cmd *cobra.Command, args []string) error {
		var err error
		if len(from) == 0 {
			return errors.New("the from option is required")
		}
		fromTime, err = time.Parse(time.RFC3339, from)
		if err != nil {
			return fmt.Errorf("%v is not a valid date in the YYYY-MM-DDTHH:mm:SS+ZZ:ZZ format", from)
		}
		if len(to) == 0 {
			return errors.New("the to option is required")
		}
		toTime, err = time.Parse(time.RFC3339, to)
		if err != nil {
			return fmt.Errorf("%v is not a valid date in the YYYY-MM-DDTHH:mm:SS+ZZ:ZZ format", to)
		}
		return nil
	},
	RunE: cache,
}

func cache(cmd *cobra.Command, args []string) error {
	ds, err := dal.New("filesystem", objects.BackendOptions{Path: "../../data/", From: fromTime, To: toTime})
	if err != nil {
		return err
	}
	cs := ds.CacheStats()
	spew.Dump(cs)
	return nil
}
