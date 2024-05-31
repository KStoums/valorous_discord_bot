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

func UnmuteCommand() commands.SlashCommand {
	return commands.SlashCommand{
		Name: "unmute",
		ArgsFunc: commands.ArgsFromStructs(
			commands.SlashCommandArg{
				Name:        "member",
				Type:        discordgo.ApplicationCommandOptionUser,
				Description: i18n.Get(discordgo.French, "moderation_commands.unmute.args.member"),
				Required:    true,
			},
		),
		Enabled:     true,
		Description: i18n.Get(discordgo.French, "moderation_commands.unmute.unmute_description"),
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

			target := i.ApplicationCommandData().Options[0].UserValue(s)
			if target.ID == s.State.User.ID {
				err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Embeds: embed.New().
							SetTitle(i18n.Get(discordgo.French, "moderation_commands.errors.title")).
							SetDescription(i18n.Get(discordgo.French, "moderation_commands.errors.unmute_member_is_bot")).
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

			guildMember, err := s.GuildMember(config.ConfigInstance.GuildId, target.ID)
			if err != nil {
				log.Logger.Error(err)
				return
			}

			if strutils.ContainString(guildMember.Roles, config.ConfigInstance.Roles.MutedRole) {
				err = s.GuildMemberRoleRemove(config.ConfigInstance.GuildId, target.ID, config.ConfigInstance.Roles.MutedRole)
				if err != nil {
					log.Logger.Error(err)
					return
				}

				err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Embeds: embed.New().
							SetTitle(i18n.Get(discordgo.French, "moderation_commands.unmute.unmute_response_title")).
							SetDescription(i18n.Get(discordgo.French, "moderation_commands.unmute.unmute_response_description",
								i18n.Vars{
									"memberMention": target.Mention(),
								})).
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

				//Logs
				err = utils.SendLogToDiscordLogChannel(s, i18n.Get(discordgo.French, "moderation_commands.logs.unmute_description",
					i18n.Vars{
						"memberMention": i.Member.Mention(),
						"mutedMember":   target.Mention(),
					}))
				if err != nil {
					log.Logger.Error(err)
					return
				}
				return
			}

			err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: embed.New().
						SetTitle(i18n.Get(discordgo.French, "moderation_commands.errors.title")).
						SetDescription(i18n.Get(discordgo.French, "moderation_commands.errors.already_unmute")).
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
		},
	}
}
