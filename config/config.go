package config

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

type App struct {
	Address string
	Static	string
	Log		string
}

type Database struct {
	Driver	string
	Address string
	Database string
	User	string
	Password string
}

type Configuration struct {
	App App
	Db	Database
}

var config *Configuration
var onece sync.Once

func LoadConfig() *Configuration{
	onece.Do(func() {
		file, err := os.Open("config.json")
		if err != nil{
			log.Fatalln("Cannot open config file", err)
		}
		decoder := json.NewDecoder(file)
		config = &Configuration{}
		err = decoder.Decode(config)
		if err != nil{
			log.Fatalln("Cannot get configuration from file", err)
		}
	})
	return config
}