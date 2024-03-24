package config

import (
	"errors"
	"github.com/goccy/go-json"
	"github.com/goroutine/template/log"
	"os"
	"path/filepath"
)

// ConfigInstance Instance of config
var ConfigInstance Config
var filePath = "./config/config.json"

// LoadConfig Load config from file
func LoadConfig() {
	file, err := os.OpenFile(filePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			err = os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
			if err != nil {
				log.Logger.Fatal("load_jsonfile_1:", err)
				return
			}

			log.Logger.Warn(filePath + " not found, creation in progress...")
			configFile, err := os.Create(filePath)
			if err != nil {
				log.Logger.Fatal("load_jsonfile_2:", err)
				return
			}
			defer configFile.Close()

			encoder := json.NewEncoder(configFile)
			encoder.SetIndent(" ", "  ")
			err = encoder.Encode(&Config{})
			if err != nil {
				log.Logger.Fatal("load_jsonfile_3:", err)
				return
			}

			log.Logger.Info(filePath + " created")
			return
		}
		log.Logger.Fatal("load_jsonfile_3:", err)
		return
	}

	err = json.NewDecoder(file).Decode(&ConfigInstance)
	if err != nil {
		log.Logger.Fatal("load_jsonfile_4:", err)
		return
	}
	file.Close()

	log.Logger.Info(filePath + " loaded")
}
