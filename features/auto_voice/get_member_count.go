package auto_voice

import "github.com/bwmarrin/discordgo"

func GetMembersCountInChannel(s *discordgo.Session, guildId, channelId string) (int, error) {
	guild, err := s.State.Guild(guildId)
	if err != nil {
		return 0, err
	}

	var nb int
	for _, voiceState := range guild.VoiceStates {
		if voiceState.ChannelID == channelId {
			nb++
		}
	}

	return nb, nil
}
