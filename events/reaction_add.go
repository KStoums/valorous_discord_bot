package events

import (
	"github.com/bwmarrin/discordgo"
	"github.com/goroutine/template/config"
	"github.com/goroutine/template/features/auto_role"
)

func ReactionAddEvent(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	if r.Member.User.ID == s.State.User.ID {
		return
	}

	if r.ChannelID != config.ConfigInstance.Channels.AutoRoleRankedChannel {
		return
	}

	auto_role.AddAutoRoleRankedFeature(s, r)
}
