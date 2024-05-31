package valorant

import (
	"github.com/bwmarrin/discordgo"
	"github.com/goccy/go-json"
	"github.com/goroutine/template/commands"
	"github.com/goroutine/template/config"
	"github.com/goroutine/template/log"
	"github.com/goroutine/template/models"
	"github.com/goroutine/template/utils/embed"
	i18n "github.com/kaysoro/discordgo-i18n"
	"net/http"
)

const valorantAgentApiUrl = "https://valorant-api.com/v1/agents?language=fr-FR"

var agentCommandArgs []*discordgo.ApplicationCommandOptionChoice

type responseBodyAgent struct {
	Status int32          `json:"status"`
	Data   []models.Agent `json:"data"`
}

func init() {
	agent, err := requestAgentApi()
	if err != nil {
		panic(err)
	}

	for _, a := range agent.Data {
		agentCommandArgs = append(agentCommandArgs, &discordgo.ApplicationCommandOptionChoice{
			Name:  a.DisplayName,
			Value: a.DisplayName,
		})
	}
}

func AgentCommand() commands.SlashCommand {
	return commands.SlashCommand{
		Name:        "agent",
		Description: i18n.Get(discordgo.French, "agent.command_description"),
		Enabled:     true,
		ArgsFunc: commands.ArgsFromStructs(
			commands.SlashCommandArg{
				Name:        "agent_name",
				Description: i18n.Get(discordgo.French, "agent.agent_args"),
				Type:        discordgo.ApplicationCommandOptionString,
				Required:    true,
				Choices:     agentCommandArgs,
			},
		),
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
							SetThumbnail("https://zupimages.net/up/24/22/n6xx.png").
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

			agentName := i.ApplicationCommandData().Options[0].StringValue()

			agents, err := requestAgentApi()
			if err != nil {
				log.Logger.Error(err)
				return
			}

			var agent models.Agent
			for _, a := range agents.Data {
				if agentName == a.DisplayName {
					agent = a
				}
			}

			finalEmbeds := embed.New().
				SetTitle(i18n.Get(discordgo.French, "agent.agent_title", i18n.Vars{
					"agentName": agentName,
					"roleName":  agent.Role.DisplayName,
				})).
				SetDescription(agent.Description).
				SetColor(embed.VALOROUS).
				SetCurrentTimestamp().
				SetDefaultFooter().
				SetThumbnail(agent.DisplayIcon).
				ToMessageEmbeds()

			for _, ability := range agent.Abilities {
				finalEmbeds[0].Fields = append(finalEmbeds[0].Fields, &discordgo.MessageEmbedField{
					Name:   "💥 " + ability.DisplayName + " (" + ability.Slot + ")",
					Value:  ability.Description,
					Inline: true,
				})
			}

			err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: finalEmbeds,
				},
			})
			if err != nil {
				log.Logger.Error(err)
				return
			}
		},
	}
}

func requestAgentApi() (responseBodyAgent, error) {
	resp, err := http.Get(valorantAgentApiUrl)
	if err != nil {
		log.Logger.Error(err)
		return responseBodyAgent{}, err
	}
	defer resp.Body.Close()

	var response responseBodyAgent
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Logger.Error(err)
		return responseBodyAgent{}, err
	}

	return response, nil
}
