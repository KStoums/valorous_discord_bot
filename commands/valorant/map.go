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
	"strings"
)

const valorantApiMapUrl = "https://valorant-api.com/v1/maps?language=fr-FR"

type responseBodyMap struct {
	Status int32        `json:"status"`
	Data   []models.Map `json:"data"`
}

var mapsCommandArgs []*discordgo.ApplicationCommandOptionChoice

func init() {
	maps, err := requestMapApi()
	if err != nil {
		panic(err)
	}

	for _, m := range maps.Data {
		mapsCommandArgs = append(mapsCommandArgs, &discordgo.ApplicationCommandOptionChoice{
			Name:  m.DisplayName,
			Value: m.DisplayName,
		})
	}
}

func MapCommand() commands.SlashCommand {
	return commands.SlashCommand{
		Name:        "map",
		Description: i18n.Get(discordgo.French, "map.command_description"),
		Enabled:     true,
		ArgsFunc: commands.ArgsFromStructs(
			commands.SlashCommandArg{
				Name:        "map_name",
				Description: i18n.Get(discordgo.French, "map.map_args"),
				Type:        discordgo.ApplicationCommandOptionString,
				Required:    true,
				Choices:     mapsCommandArgs,
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

			mapName := i.ApplicationCommandData().Options[0].StringValue()

			maps, err := requestMapApi()
			if err != nil {
				log.Logger.Error(err)
				return
			}

			var mapp models.Map
			for _, m := range maps.Data {
				if m.DisplayName == mapName {
					mapp = m
					break
				}
			}

			err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: embed.New().
						SetTitle(i18n.Get(discordgo.French, "map.map_title", i18n.Vars{
							"mapName": mapName,
						})).
						SetDescription(mapp.NarrativeDescription).
						SetColor(embed.VALOROUS).
						SetCurrentTimestamp().
						SetDefaultFooter().
						SetThumbnail(mapp.DisplayIcon).
						SetImage(mapp.Splash).
						AddInlinedField("💣 Sites", strings.Split(mapp.TacticalDescription, " ")[0]).
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

func requestMapApi() (responseBodyMap, error) {
	resp, err := http.Get(valorantApiMapUrl)
	if err != nil {
		log.Logger.Error(err)
		return responseBodyMap{}, err
	}
	defer resp.Body.Close()

	var response responseBodyMap
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Logger.Error(err)
		return responseBodyMap{}, err
	}

	return response, nil
}
