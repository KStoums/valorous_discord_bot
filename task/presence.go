package task

import (
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/goroutine/template/config"
	"github.com/goroutine/template/log"
)

var i int

type BotPresence struct {
	s *discordgo.Session
}

func NewBotPresence(s *discordgo.Session) *BotPresence {
	return &BotPresence{s: s}
}

func (b BotPresence) Run() {
	members, err := b.s.GuildMembers(config.ConfigInstance.GuildId, "", 100)
	if err != nil {
		log.Logger.Error(errors.New("unable to get guild members : " + err.Error()))
		return
	}
	presenceName := []string{fmt.Sprintf("%d/100 membres", len(members)), "Valorant"}

	if i == 0 {
		err = b.s.UpdateWatchStatus(1, presenceName[i])
		if err != nil {
			log.Logger.Error(errors.New("unable to update bot status : " + err.Error()))
			return
		}
	} else {
		err = b.s.UpdateGameStatus(1, presenceName[i])
		if err != nil {
			log.Logger.Error(errors.New("unable to update bot status : " + err.Error()))
			return
		}
	}

	i++
	if i == len(presenceName) {
		i = 0
	}
}

func (b BotPresence) CronString() string {
	return "@every 10s"
}

func (b BotPresence) Name() string {
	return "BotPresence"
}

func (b BotPresence) RunOnStart() bool {
	return true
}
