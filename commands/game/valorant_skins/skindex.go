package valorant_skins

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/bwmarrin/discordgo"
	"github.com/goroutine/template/commands"
	"github.com/goroutine/template/config"
	"github.com/goroutine/template/log"
	"github.com/goroutine/template/models"
	"github.com/goroutine/template/utils/embed"
	i18n "github.com/kaysoro/discordgo-i18n"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

const valorantApiWeaponSkinsUrl = "https://valorant-api.com/v1/weapons/skins"

type responseBodySkins struct {
	Status int32               `json:"status"`
	Data   []models.WeaponSkin `json:"data"`
}

var valorantSkinCount int

func init() {
	skins, err := requestSkinApi()
	if err != nil {
		panic(err)
	}

	valorantSkinCount = len(skins.Data)
}

func SkinDexCommand(database *mongo.Database) commands.SlashCommand {
	return commands.SlashCommand{
		Name:        "skindex",
		Description: i18n.Get(discordgo.French, "game_valorant_skins.skindex_description"),
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

			ctx := context.Background()
			collection := database.Collection("game_valorant_skins")

			var userData *models.GameValorantSkins
			err := collection.FindOne(ctx, models.GameValorantSkins{UserId: i.Member.User.ID}).Decode(&userData)
			if err != nil {
				if errors.Is(err, mongo.ErrNoDocuments) {
					newUser := &models.GameValorantSkins{UserId: i.Member.User.ID}

					_, err = collection.InsertOne(ctx, newUser)
					if err != nil {
						log.Logger.Error(err)
						return
					}

					userData = newUser
				} else {
					log.Logger.Error(err)
					return
				}
			}

			err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: embed.New().
						SetTitle(i18n.Get(discordgo.French, "game_valorant_skins.skindex_response_title", i18n.Vars{
							"skindexAuthorUsername": i.Member.User.GlobalName,
						})).
						SetDescription(i18n.Get(discordgo.French, "game_valorant_skins.skindex_response_description", i18n.Vars{
							"skinCount":         len(userData.SkinsCollectedName),
							"valorantSkinCount": valorantSkinCount,
						})).
						SetCurrentTimestamp().
						SetDefaultFooter().
						SetColor(embed.VALOROUS).
						SetThumbnail("https://zupimages.net/up/24/16/yte5.png").
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

func requestSkinApi() (responseBodySkins, error) {
	resp, err := http.Get(valorantApiWeaponSkinsUrl)
	if err != nil {
		log.Logger.Error(err)
		return responseBodySkins{}, err
	}
	defer resp.Body.Close()

	var response responseBodySkins
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Logger.Error(err)
		return responseBodySkins{}, err
	}

	return response, nil
}
