package session

import "github.com/spf13/viper"

func init() {
	viper.SetDefault("session_inactivity_timeout", 1800)
}
