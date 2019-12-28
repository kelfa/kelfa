package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.kelfa.io/pkg/dal"
	"go.kelfa.io/pkg/dal/objects"
)

func init() {
	rootCmd.AddCommand(dbCmd)
	dbCmd.AddCommand(dbImportCmd)
}

var dbCmd = &cobra.Command{
	Use:   "database",
	Short: "see hits rates",
}

var dbImportCmd = &cobra.Command{
	Use:   "import",
	Short: "import data",
	RunE:  dbImport,
}

func dbImport(cmd *cobra.Command, args []string) error {
	fs, err := dal.NewDataSource("filesystem", objects.BackendOptions{Path: viper.GetString("data_folder")})
	if err != nil {
		return err
	}
	sql, err := dal.NewWritableDataSource("sqlite", objects.BackendOptions{Path: fmt.Sprintf("%s/.kelfa/database.db", os.Getenv("HOME"))})
	if err != nil {
		return err
	}
	fromTime, err := fs.DataBeginTime()
	if err != nil {
		return err
	}
	toTime, err := fs.DataEndTime()
	if err != nil {
		return err
	}
	dps, err := fs.GetDataPoints(*fromTime, *toTime)
	if err != nil {
		return err
	}
	dpssql, err := sql.GetDataPoints(*fromTime, *toTime)
	if err != nil {
		return err
	}
	for _, dp := range dps {
		var present bool
		for _, dpsql := range dpssql {
			if dp.ID == dpsql.ID {
				present = true
				break
			}
		}
		if !present {
			fmt.Printf("Going to add %s %s...", dp.DateTime, dp.ID)
			if err := sql.AddDataPoint(dp); err != nil {
				log.Fatalf("%v", err)
			}
			fmt.Printf(" done\n")
		} else {
			fmt.Printf("Not going to add %s %s. Already present\n", dp.DateTime, dp.ID)
		}
	}
	return nil
}
