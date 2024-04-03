package utils

import (
	"github.com/bwmarrin/discordgo"
	"github.com/goroutine/template/config"
	"github.com/goroutine/template/utils/embed"
)

func SendLogToDiscordLogChannel(s *discordgo.Session, description string) error {
	_, err := s.ChannelMessageSendEmbed(config.ConfigInstance.Channels.LogChannel, embed.New().
		SetTitle("Logs").
		SetColor(embed.VALOROUS).
		SetCurrentTimestamp().
		SetDefaultFooter().
		SetThumbnail("https://zupimages.net/up/24/14/lweg.png").
		SetDescription(description).
		ToMessageEmbed())
	return err
}
