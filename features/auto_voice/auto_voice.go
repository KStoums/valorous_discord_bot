package auto_voice

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/goroutine/template/log"
)

func AutoVoiceFeature(s *discordgo.Session, v *discordgo.VoiceStateUpdate, categoryId string) {
	autoVoiceChannelCategory, err := s.Channel(categoryId)
	if err != nil {
		log.Logger.Error(err)
		return
	}

	member, err := s.GuildMember(v.GuildID, v.Member.User.ID)
	if err != nil {
		log.Logger.Error(err)
		return
	}

	autoVoiceChannel, err := s.GuildChannelCreateComplex(v.GuildID, discordgo.GuildChannelCreateData{
		Name:                 fmt.Sprintf("│🎤│ Salon de %s", member.User.GlobalName),
		Type:                 discordgo.ChannelTypeGuildVoice,
		PermissionOverwrites: autoVoiceChannelCategory.PermissionOverwrites,
		ParentID:             autoVoiceChannelCategory.ID,
		UserLimit:            6,
	})
	if err != nil {
		log.Logger.Error(err)
		return
	}

	err = s.GuildMemberMove(v.GuildID, v.UserID, &autoVoiceChannel.ID)
	if err != nil {
		log.Logger.Error(err)
		return
	}

	log.Logger.Debug("An auto voice channel has been created by " + member.User.GlobalName)
	return
}
