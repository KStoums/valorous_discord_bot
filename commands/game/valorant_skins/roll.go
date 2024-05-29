package valorant_skins

import (
	"github.com/bwmarrin/discordgo"
	"github.com/goroutine/template/commands"
	"github.com/goroutine/template/config"
	"github.com/goroutine/template/log"
	"github.com/goroutine/template/utils/embed"
	i18n "github.com/kaysoro/discordgo-i18n"
	"math/rand"
)

func GameValorantSkins() commands.SlashCommand {
	return commands.SlashCommand{
		Name:        "roll",
		Description: i18n.Get(discordgo.French, "game_valorant_skins.roll_description"),
		Enabled:     true,
		Handler: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			if i.ChannelID != config.ConfigInstance.Channels.BotCommand && i.ChannelID != config.ConfigInstance.Channels.BotCommandAdmin {
				err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Embeds: embed.New().
							SetTitle(i18n.Get(discordgo.French, "game_valorant_skins.errors.title")).
							SetDescription(i18n.Get(discordgo.French, "game_valorant_skins.errors.not_in_bot_command_channel_description")).
							SetCurrentTimestamp().
							SetDefaultFooter().
							SetColor(embed.VALOROUS).
							ToMessageEmbeds(),
						Flags: discordgo.MessageFlagsEphemeral,
					},
				})
				if err != nil {
					log.Logger.Error(err)
					return
				}
				return
			}

			skins, err := requestSkinApi()
			if err != nil {
				log.Logger.Error(err)
				return
			}

			randomInt := rand.Intn(len(skins.Data))
			err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: embed.New().
						SetTitle(i18n.Get(discordgo.French, "game_valorant_skins.roll_response_title")).
						SetColor(embed.VALOROUS).
						SetCurrentTimestamp().
						SetDefaultFooter().
						AddInlinedField("ℹ️ Nom du skin", skins.Data[randomInt].DisplayName).
						SetImage(skins.Data[randomInt].DisplayIcon).
						ToMessageEmbeds(),
				},
			})
			if err != nil {
				log.Logger.Error(err)
				return
			}
		},
	}
}
