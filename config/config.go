package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

const confVer = "1.0.0"

const confFileName = "conf.json"

type ConfType struct {
	Version        string              `json:"version"`
	Debug          bool                `json:"debug"`
	SyncRepository []map[string]string `json:"sync_repository"`
}

func LoadConfig() (Conf ConfType) {

	currentPath, _ := os.Getwd()

	fmt.Println(filepath.Join(currentPath, confFileName))

	data, _ := ioutil.ReadFile(filepath.Join(currentPath, confFileName))

	json.Unmarshal(data, &Conf)

	return
}
