package config

import "testProject/pkg/repository"

var dbConfig = repository.Config{
	Host:     "localhost",
	Port:     "5432",
	Username: "postgres",
	Password: "1234",
	DBName:   "test_task",
	SSLMode:  "disable",
}

func GetDbConfig() (repository.Config, error) {
	return dbConfig, nil
}
