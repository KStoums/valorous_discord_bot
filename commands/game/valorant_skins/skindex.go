package valorant_skins

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/goroutine/template/commands"
	"github.com/goroutine/template/config"
	"github.com/goroutine/template/database"
	"github.com/goroutine/template/log"
	"github.com/goroutine/template/models"
	"github.com/goroutine/template/utils/embed"
	i18n "github.com/kaysoro/discordgo-i18n"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"os"
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

func SkinDexCommand() commands.SlashCommand {
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
			collection := database.MongoDb.Collection(os.Getenv("GAME_VALORANT_SKINS_COLLECTION_NAME"))

			var userData *models.UserSkinGame
			err := collection.FindOne(ctx, bson.D{{"_id", i.Member.User.ID}}).Decode(&userData)
			if err != nil {
				if errors.Is(err, mongo.ErrNoDocuments) {
					newUser := &models.UserSkinGame{UserId: i.Member.User.ID, RollRemaining: 5}

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
						AddInlinedField("📦 Skins collectés", fmt.Sprintf("%d/%d", len(userData.SkinsCollectedName), valorantSkinCount)).
						SetCurrentTimestamp().
						SetDefaultFooter().
						SetColor(embed.VALOROUS).
						SetThumbnail("https://zupimages.net/up/24/22/b5js.png").
						SetImage("https://admin.esports.gg/wp-content/uploads/2023/12/VALORANT-Overdrive-bundle.jpg").
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
