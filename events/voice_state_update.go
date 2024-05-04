package events

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/goroutine/template/config"
	"github.com/goroutine/template/features/auto_voice"
	"github.com/goroutine/template/log"
	"strings"
)

func VoiceStateUpdateEvent(s *discordgo.Session, v *discordgo.VoiceStateUpdate) {
	switch v.ChannelID {
	case config.ConfigInstance.Channels.AutoVoiceChannel, config.ConfigInstance.Channels.AutoVoiceChannelTeam, config.ConfigInstance.Channels.AutoVoiceChannelAdmin:
		autoVoiceChannel, err := s.Channel(v.ChannelID)
		if err != nil {
			log.Logger.Error(err)
			return
		}

		auto_voice.AutoVoiceFeature(s, v, autoVoiceChannel.ParentID)
		return

	default:
		beforeChannel, err := s.Channel(v.BeforeUpdate.ChannelID)
		if err != nil {
			log.Logger.Error(err)
			return
		}

		if beforeChannel.ParentID != config.ConfigInstance.Channels.AutoVoiceCategory && beforeChannel.ParentID != config.ConfigInstance.Channels.TeamCategory &&
			beforeChannel.ParentID != config.ConfigInstance.Channels.AdminCategory {
			return
		}

		if strings.Contains(beforeChannel.Name, "🎤") {
			membersCountInChannel, err := auto_voice.GetMembersCountInChannel(s, v.GuildID, v.BeforeUpdate.ChannelID)
			if membersCountInChannel == 0 {
				_, err = s.ChannelDelete(beforeChannel.ID)
				if err != nil {
					log.Logger.Error(err)
					return
				}

				log.Logger.Debug(fmt.Sprintf("The vocal salon %s has been deleted because the number of members is equal to 0", beforeChannel.Name))
				return
			}
		}
	}
}
