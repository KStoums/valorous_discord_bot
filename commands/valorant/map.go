package valorant

import (
	"github.com/bwmarrin/discordgo"
	"github.com/goccy/go-json"
	"github.com/goroutine/template/commands"
	"github.com/goroutine/template/log"
	"github.com/goroutine/template/models"
	"github.com/goroutine/template/utils/embed"
	i18n "github.com/kaysoro/discordgo-i18n"
	"net/http"
	"strings"
)

const valorantApiMapUrl = "https://valorant-api.com/v1/maps"

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
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "Ascent",
						Value: "Ascent",
					},
					{
						Name:  "Bind",
						Value: "Bind",
					},
					{
						Name:  "Fracture",
						Value: "Fracture",
					},
					{
						Name:  "Haven",
						Value: "Haven",
					},
					{
						Name:  "Icebox",
						Value: "Icebox",
					},
					{
						Name:  "Lotus",
						Value: "Lotus",
					},
					{
						Name:  "Pearl",
						Value: "Pearl",
					},
					{
						Name:  "Split",
						Value: "Split",
					},
					{
						Name:  "Sunset",
						Value: "Sunset",
					},
				},
			},
		),
		Handler: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			mapName := i.ApplicationCommandData().Options[0].StringValue()

			resp, err := http.Get(valorantApiMapUrl)
			if err != nil {
				log.Logger.Error(err)
				return
			}
			defer resp.Body.Close()

			var responseBody struct {
				Status int32        `json:"status"`
				Data   []models.Map `json:"data"`
			}
			err = json.NewDecoder(resp.Body).Decode(&responseBody)
			if err != nil {
				log.Logger.Error(err)
				return
			}

			var mapp models.Map
			for _, m := range responseBody.Data {
				if m.DisplayName == mapName {
					mapp = m
					break
				}
			}

			finalEmbeds := embed.New().
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
				ToMessageEmbeds()

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