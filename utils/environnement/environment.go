package environnement

import (
	"github.com/goroutine/template/log"
	"os"
)

func GetToken() string {
	return os.Getenv("DISCORD_TOKEN")
}

func CheckEnvs() {
	if GetToken() == "" {
		log.Logger.Fatal("DISCORD_TOKEN is not set")
	}
}
