package middleware

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

type StoredConfig struct {
	Type string
	Config string
}

const(
	lastConfigStorePath = "/home/ubuntu//store/lighting/last/"
	typeFile = "type.data"
	configFile = "config.data"
)

func GetLastConfig() *StoredConfig {
	dat, err := ioutil.ReadFile(lastConfigStorePath + typeFile)
	if err != nil {
		return &StoredConfig{
			Type: "one",
			Config: "ffffff",
		}
	}

	conf, err := ioutil.ReadFile(lastConfigStorePath + configFile)
	if err != nil {
		return &StoredConfig{
			Type: "one",
			Config: "ffffff",
		}
	}

	return &StoredConfig{
		Type: strings.ToLower(string(dat)),
		Config: string(conf),
	}
}

func StoreLastConfig(confType string, config interface{}) {
	err := ioutil.WriteFile(lastConfigStorePath + typeFile, []byte(confType), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	b, err := json.Marshal(config)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ioutil.WriteFile(lastConfigStorePath + configFile, []byte(b), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}