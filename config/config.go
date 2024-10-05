package config

import (
	"encoding/json"
	"fmt"
	"os"
)

var Token string
var Botprefix string
var config *Config

type Config struct {
	Token     string `json:"token"`
	BotPreFix string `json:"botPrefix"`
}

// ReadConfig reads the config.json file and unmarshals into the Config Struct
func ReadConfig() error {
	fmt.Println("Reading config.json...")
	file, err := os.ReadFile("./config.json")

	if err != nil {
		return err
	}

	fmt.Println("Unmarshalling config.json...")

	//unmarshall file into config struct
	err = json.Unmarshal(file, &config)

	if err != nil {
		fmt.Println("Error unmarshalling config.json")
		return err
	}

	Token = config.Token
	Botprefix = config.BotPreFix

	return nil
}
