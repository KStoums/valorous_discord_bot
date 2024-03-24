package environnement

import (
	"github.com/goroutine/template/log"
	"os"
)

func GetToken() string {
	return os.Getenv("DISCORD_TOKEN")
}

func GetMariaDsn() string {
	return os.Getenv("MARIADB_DSN")
}

func CheckEnvs() {
	if GetToken() == "" {
		log.Logger.Fatal("DISCORD_TOKEN is not set")
	}
	if GetMariaDsn() == "" {
		log.Logger.Fatal("MARIADB_DSN is not set")
	}
}
