package moderation

import (
	"github.com/bwmarrin/discordgo"
	"github.com/goroutine/template/commands"
	"github.com/goroutine/template/config"
	"github.com/goroutine/template/log"
	"github.com/goroutine/template/utils"
	"github.com/goroutine/template/utils/embed"
	"github.com/goroutine/template/utils/strutils"
	i18n "github.com/kaysoro/discordgo-i18n"
)

func ClearCommand() commands.SlashCommand {
	return commands.SlashCommand{
		Name: "clear",
		ArgsFunc: commands.ArgsFromStructs(
			commands.SlashCommandArg{
				Name:        "message_count",
				Type:        discordgo.ApplicationCommandOptionInteger,
				Description: i18n.Get(discordgo.French, "moderation_commands.clear.clear_message_count_args"),
				Required:    true,
			},
		),
		Enabled:     true,
		Description: i18n.Get(discordgo.French, "moderation_commands.clear.clear_command_description"),
		Handler: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			if !strutils.ContainString(i.Member.Roles, config.ConfigInstance.Roles.AdministrationRole) {
				err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Embeds: embed.New().
							SetTitle(i18n.Get(discordgo.French, "moderation_commands.errors.title")).
							SetDescription(i18n.Get(discordgo.French, "moderation_commands.errors.description")).
							SetColor(embed.VALOROUS).
							SetCurrentTimestamp().
							SetDefaultFooter().
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

			messageCountToDelete := i.ApplicationCommandData().Options[0].IntValue()
			channelMessages, err := s.ChannelMessages(i.ChannelID, 100, "", "", "")
			if err != nil {
				log.Logger.Error(err)
				return
			}

			var messageList []string
			for y, message := range channelMessages {
				if int64(y) < messageCountToDelete {
					y++
					messageList = append(messageList, message.ID)
				}
			}

			err = s.ChannelMessagesBulkDelete(i.ChannelID, messageList)
			if err != nil {
				log.Logger.Error(err)
				return
			}

			err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: embed.New().
						SetTitle(i18n.Get(discordgo.French, "moderation_commands.clear.clear_response_title")).
						SetDescription(i18n.Get(discordgo.French, "moderation_commands.clear.clear_response_description",
							i18n.Vars{
								"messageCount": messageCountToDelete,
							})).
						SetDefaultFooter().
						SetCurrentTimestamp().
						SetColor(embed.VALOROUS).
						ToMessageEmbeds(),
					Flags: discordgo.MessageFlagsEphemeral,
				},
			})
			if err != nil {
				log.Logger.Error(err)
				return
			}

			//Logs
			channel, err := s.Channel(i.ChannelID)
			if err != nil {
				log.Logger.Error(err)
				return
			}

			err = utils.SendLogToDiscordLogChannel(s, i18n.Get(discordgo.French, "moderation_commands.logs.clear_message_description",
				i18n.Vars{
					"memberMention":  i.Member.Mention(),
					"messageCount":   messageCountToDelete,
					"channelMention": channel.Mention(),
				}))
			if err != nil {
				log.Logger.Error(err)
				return
			}
		},
	}
}
