package auto_voice

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
)

func AutoVoiceFeature(s *discordgo.Session, v *discordgo.VoiceStateUpdate, categoryId string) {
	autoVoiceChannelCategory, err := s.Channel(categoryId)
	if err != nil {
		log.Error().Err(err)
		return
	}

	member, err := s.GuildMember(v.GuildID, v.Member.User.ID)
	if err != nil {
		log.Logger.Err(err)
		return
	}

	autoVoiceChannel, err := s.GuildChannelCreateComplex(v.GuildID, discordgo.GuildChannelCreateData{
		Name:                 fmt.Sprintf("│🎤│ Salon de %s", member.User.GlobalName),
		Type:                 discordgo.ChannelTypeGuildVoice,
		PermissionOverwrites: autoVoiceChannelCategory.PermissionOverwrites,
		ParentID:             autoVoiceChannelCategory.ID,
	})
	if err != nil {
		log.Error().Err(err)
		return
	}

	err = s.GuildMemberMove(v.GuildID, v.UserID, &autoVoiceChannel.ID)
	if err != nil {
		log.Error().Err(err)
		return
	}

	log.Debug().Msg(fmt.Sprintf("An auto voice channel has been created by %s", member.User.GlobalName))
	return
}
