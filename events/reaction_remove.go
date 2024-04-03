package events

import (
	"github.com/bwmarrin/discordgo"
	"github.com/goroutine/template/config"
	"github.com/goroutine/template/features/auto_role"
)

func ReactionRemoveEvent(s *discordgo.Session, r *discordgo.MessageReactionRemove) {
	if r.UserID == s.State.User.ID {
		return
	}

	if r.ChannelID != config.ConfigInstance.Channels.AutoRoleRankedChannel {
		return
	}

	auto_role.RemoveAutoRoleRankedFeature(s, r)
}
