package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.kelfa.io/pkg/dal"
	"go.kelfa.io/pkg/dal/objects"
	"go.kelfa.io/pkg/session"
)

func init() {
	rootCmd.AddCommand(dbCmd)
	dbCmd.AddCommand(dbImportCmd)
	dbImportCmd.PersistentFlags().Bool("full", false, "rescan all files")
	viper.BindPFlag("full", dbImportCmd.PersistentFlags().Lookup("full"))
	dbCmd.AddCommand(dbSessionsUpdateCmd)
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
	if viper.GetBool("full") {
		return dbDeepImport(cmd, args)
	} else {
		return dbIncrementalImport(cmd, args)
	}
}

func dbIncrementalImport(cmd *cobra.Command, args []string) error {
	fs, err := dal.NewDataSource("filesystem", objects.BackendOptions{Path: viper.GetString("data_folder")})
	if err != nil {
		return err
	}
	sql, err := dal.NewWritableDataSource("sqlite", objects.BackendOptions{Path: fmt.Sprintf("%s/.kelfa/database.db", os.Getenv("HOME"))})
	if err != nil {
		return err
	}
	fromTime, err := sql.DataEndTime()
	if err != nil {
		return err
	}
	toTime, err := fs.DataEndTime()
	if err != nil {
		return err
	}
	dpssql, err := sql.GetDataPoints(*fromTime, *toTime)
	if err != nil {
		return err
	}
	dps, err := fs.GetDataPoints(*fromTime, *toTime)
	if err != nil {
		return err
	}
	var imported time.Time
	for _, dp := range dps {
		var present bool
		for _, dpsql := range dpssql {
			if dp.ID == dpsql.ID {
				present = true
				break
			}
		}
		if imported.Format("20060102") != dp.DateTime.Format("20060102") {
			fmt.Printf("Importing data of %s\n", dp.DateTime.Format("02/01/2006"))
			imported = dp.DateTime
		}
		if !present {
			if err := sql.AddDataPoint(dp); err != nil {
				log.Fatalf("%v", err)
			}
		}
	}
	return nil
}

func dbDeepImport(cmd *cobra.Command, args []string) error {
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
	var imported time.Time
	for _, dp := range dps {
		var present bool
		for _, dpsql := range dpssql {
			if dp.ID == dpsql.ID {
				present = true
				break
			}
		}
		if imported.Format("20060102") != dp.DateTime.Format("20060102") {
			fmt.Printf("Importing data of %s\n", dp.DateTime.Format("02/01/2006"))
			imported = dp.DateTime
		}
		if !present {
			if err := sql.AddDataPoint(dp); err != nil {
				log.Fatalf("%v", err)
			}
		}
	}
	return nil
}

var dbSessionsUpdateCmd = &cobra.Command{
	Use:   "sessions-update",
	Short: "update sessions data",
	RunE:  dbSessionUpdate,
}

func dbSessionUpdate(cmd *cobra.Command, args []string) error {
	fromTime, err := ds.DataBeginTime()
	if err != nil {
		return err
	}
	toTime, err := ds.DataEndTime()
	if err != nil {
		return err
	}
	dsdps, err := ds.GetDataPoints(*fromTime, *toTime)
	if err != nil {
		return err
	}
	var ss session.Sessions
	for _, s := range dsdps {
		s := s
		ss.AddDataPoint(&s)
	}
	ss.SplitSessions(viper.GetDuration("session_inactivity_timeout"))
	err = ds.AddSessions(ss)
	return err
}
