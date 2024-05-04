package task

import (
	"errors"
	"github.com/bwmarrin/discordgo"
	"github.com/goroutine/template/api"
	"github.com/goroutine/template/config"
	"github.com/goroutine/template/log"
	"strconv"
)

var _ api.Task = (*GuildMemberCountTask)(nil)

type GuildMemberCountTask struct {
	s *discordgo.Session
}

func NewGuildMemberCountTask(s *discordgo.Session) *GuildMemberCountTask {
	return &GuildMemberCountTask{s: s}
}

func (g *GuildMemberCountTask) CronString() string {
	return "@every 10m"
}

func (g *GuildMemberCountTask) Name() string {
	return "GuildMemberCountTask"
}

func (g *GuildMemberCountTask) Run() {
	members, err := g.s.GuildMembers(config.ConfigInstance.GuildId, "", 100)
	if err != nil {
		log.Logger.Error(errors.New("unable to get guild members : " + err.Error()))
		return
	}

	if _, err = g.s.ChannelEdit(config.ConfigInstance.Channels.MemberCount, &discordgo.ChannelEdit{
		Name: "│👥│Membres: " + strconv.Itoa(len(members)) + "/100",
	}); err != nil {
		log.Logger.Error("Could not edit member count channel: ", err)
	}
}
