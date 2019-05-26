package config

import (
	"encoding/json"
	"os"
)

type BasicsCon struct {
	DriverName string `json:"driverName"`
	Dsn string `json:"dsn"`
}

var (
	BasicsConfig *BasicsCon
)

func init() {
	filePath := "./InfrastructureConfig.json"
	file, e := os.Open(filePath)
	if e != nil {
		panic(e.Error())
	}

	BasicsConfig = &BasicsCon{}
	decoder := json.NewDecoder(file)
	e = decoder.Decode(BasicsConfig)
	if e != nil {
		panic(e.Error())
	}

}

