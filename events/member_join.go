package events

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/goroutine/template/config"
	"github.com/goroutine/template/log"
	"github.com/goroutine/template/utils/embed"
	i18n "github.com/kaysoro/discordgo-i18n"
)

func MemberJoinEvent(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	guildMembers, err := s.GuildMembers(m.GuildID, "", 1000)
	if err != nil {
		log.Logger.Error(err)
	}

	_, err = s.ChannelMessageSendEmbeds(config.ConfigInstance.Channels.WelcomeChannel, embed.New().
		SetTitle(i18n.Get(discordgo.French, "event.member_join_title", i18n.Vars{})+m.User.GlobalName).
		SetDescription(fmt.Sprintf("%s **%d**", i18n.Get(discordgo.French, "event.member_join_description"), len(guildMembers))).
		SetColor(embed.VALOROUS).
		SetThumbnail(m.AvatarURL("")).
		SetDefaultFooter().
		SetCurrentTimestamp().
		ToMessageEmbeds())
	if err != nil {
		log.Logger.Error(err)
		return
	}

	err = s.GuildMemberRoleAdd(m.GuildID, m.Member.User.ID, config.ConfigInstance.Roles.MemberRole)
	if err != nil {
		log.Logger.Error(err)
		return
	}
}