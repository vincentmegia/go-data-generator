package common

import "os"

type ConfigurationManager struct{}

type DBConfiguration struct {
	User     string
	Password string
	Protocol string
	Hostname string
	DBName   string
}

func (cm ConfigurationManager) GetDBConfig() *DBConfiguration {
	return &DBConfiguration{
		User:     os.Getenv("DBUser"),
		Password: os.Getenv("DBPassword"),
		Protocol: os.Getenv("DBProtocol"),
		Hostname: os.Getenv("DBHostname"),
		DBName:   os.Getenv("DBName"),
	}
}
