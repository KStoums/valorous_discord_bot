package task

import (
	"github.com/bwmarrin/discordgo"
	"github.com/goroutine/template/api"
	"github.com/goroutine/template/config"
	"github.com/goroutine/template/helpers"
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

func (g GuildMemberCountTask) CronString() string {
	return "@every 10m"
}

func (g GuildMemberCountTask) Name() string {
	return "GuildMemberCountTask"
}

func (g GuildMemberCountTask) Run() {
	membersCount, err := helpers.GetMembersCount(g.s)
	if err != nil {
		log.Logger.Error("Could not get members count: ", err)
	}

	if _, err = g.s.ChannelEdit(config.ConfigInstance.Channels.MemberCount, &discordgo.ChannelEdit{
		Name: "│👥│Membres: " + strconv.Itoa(membersCount) + "/100",
	}); err != nil {
		log.Logger.Error("Could not edit member count channel: ", err)
	}
}
