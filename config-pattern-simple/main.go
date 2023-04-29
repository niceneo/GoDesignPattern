package main

import (
	"fmt"
	"os"
)
type Config struct {
	DBHost string
	DBPort int
	DBName string
	LogLevel string
}
func LoadConfig()(*Config,error){
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	logLevel := os.Getenv("LOG_LEVEL")

	var port int
	if dbPort != ""{
		_,err := fmt.Sscanf(dbPort, "%d",&port)
		if err != nil {
			return nil,err
		}
	}
	return &Config{
		DBHost: dbHost,
		DBPort: port,
		DBName: dbName,
		LogLevel: logLevel,
	},nil
}
func mian(){
	config,err := LoadConfig()
	if err != nil {
		fmt.Println("Error loading config:",err)
		return
	}
	fmt.Println("DBHost:",config.DBHost)
	fmt.Println("DBPort:",config.DBPort)
	fmt.Println("DBName:",config.DBName)
	fmt.Println("LogLevel:",config.LogLevel)
}
