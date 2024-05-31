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

func MuteCommand() commands.SlashCommand {
	return commands.SlashCommand{
		Name: "mute",
		ArgsFunc: commands.ArgsFromStructs(
			commands.SlashCommandArg{
				Name:        "member",
				Type:        discordgo.ApplicationCommandOptionUser,
				Description: i18n.Get(discordgo.French, "moderation_commands.mute.args.member"),
				Required:    true,
			},
			commands.SlashCommandArg{
				Name:        "reason",
				Type:        discordgo.ApplicationCommandOptionString,
				Description: i18n.Get(discordgo.French, "moderation_commands.mute.args.reason"),
				Required:    true,
			},
		),
		Enabled:     true,
		Description: i18n.Get(discordgo.French, "moderation_commands.mute.mute_description"),
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
							SetThumbnail("https://zupimages.net/up/24/22/u2e8.png").
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
							SetDescription(i18n.Get(discordgo.French, "moderation_commands.errors.muted_member_is_bot")).
							SetColor(embed.VALOROUS).
							SetCurrentTimestamp().
							SetThumbnail("https://zupimages.net/up/24/22/u2e8.png").
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

			if !strutils.ContainString(guildMember.Roles, config.ConfigInstance.Roles.MutedRole) {
				err = s.GuildMemberRoleAdd(config.ConfigInstance.GuildId, target.ID, config.ConfigInstance.Roles.MutedRole)
				if err != nil {
					log.Logger.Error(err)
					return
				}

				muteReason := i.ApplicationCommandData().Options[1].StringValue()
				err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Embeds: embed.New().
							SetTitle(i18n.Get(discordgo.French, "moderation_commands.mute.muted_response_title")).
							SetDescription(i18n.Get(discordgo.French, "moderation_commands.mute.muted_response_description",
								i18n.Vars{
									"memberMention": target.Mention(),
									"muteReason":    muteReason,
								})).
							SetColor(embed.VALOROUS).
							SetCurrentTimestamp().
							SetThumbnail("https://zupimages.net/up/24/22/vr0y.png").
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
				err = utils.SendLogToDiscordLogChannel(s, i18n.Get(discordgo.French, "moderation_commands.logs.muted_description",
					i18n.Vars{
						"memberMention": i.Member.Mention(),
						"mutedMember":   target.Mention(),
						"muteReason":    muteReason,
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
						SetDescription(i18n.Get(discordgo.French, "moderation_commands.errors.already_muted")).
						SetCurrentTimestamp().
						SetThumbnail("https://zupimages.net/up/24/22/vr0y.png").
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
