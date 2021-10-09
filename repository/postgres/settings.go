package postgres

import "fmt"

type Settings struct {
	Host     string
	Port     uint64
	User     string
	Password string
	DBName   string `mapstructure:"db_name"`
	SSLMode  string `mapstructure:"ssl_mode"`
	TimeZone string `mapstructure:"time_zone"`
}

func (s Settings) DSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		s.Host, s.User, s.Password, s.DBName, s.Port, s.SSLMode, s.TimeZone)
}
