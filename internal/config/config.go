package config

import "os"

var localdbConfig = "postgresql://postgres:1234@localhost:15432/postgres?application_name=postgres&sslmode=disable"

func GetDbConfig() string {
	if s := os.Getenv("PG_DSN"); s != "" {
		return s
	}
	return localdbConfig
}
