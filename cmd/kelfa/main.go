package main

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/kelfa/kelfa/pkg/dal"
	"github.com/kelfa/kelfa/pkg/dal/objects"
)

var ds dal.WritableDataSource

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath("$HOME/.kelfa")
	viper.SetDefault("session_inactivity_timeout", "1h")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("fatal error config file: %s \n", err)
		return
	}
	ds, err = dal.NewWritableDataSource("sqlite", objects.BackendOptions{Path: fmt.Sprintf("%s/.kelfa/database.db", os.Getenv("HOME"))})
	if err != nil {
		fmt.Printf("fatal error loading datasource: %s \n", err)
		return
	}

	rootCmd.PersistentFlags().StringP("from", "f", defaultFromDate(), "from")
	rootCmd.PersistentFlags().StringP("to", "t", defaultToDate(), "to")
	rootCmd.PersistentFlags().StringP("data", "d", "", "data folder")

	if err := viper.BindPFlag("data_folder", rootCmd.PersistentFlags().Lookup("data")); err != nil {
		panic(err)
	}
	if err := viper.BindPFlag("period_from", rootCmd.PersistentFlags().Lookup("from")); err != nil {
		panic(err)
	}
	if err := viper.BindPFlag("period_to", rootCmd.PersistentFlags().Lookup("to")); err != nil {
		panic(err)
	}
}

var rootCmd = &cobra.Command{
	Use: "kelfa",
}

func defaultFromDate() string {
	dd, err := ds.DataEndTime()
	if err != nil {
		return time.Time{}.Format(time.RFC3339)
	}
	d := dd.AddDate(0, 0, -6)
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, time.UTC).Format(time.RFC3339)
}

func defaultToDate() string {
	d, err := ds.DataEndTime()
	if err != nil {
		return time.Time{}.Format(time.RFC3339)
	}
	return d.Format(time.RFC3339)
}
