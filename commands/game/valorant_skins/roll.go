package valorant_skins

import (
	"context"
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
	"math/rand"
	"os"
	"time"
)

func RollCommand() commands.SlashCommand {
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

			ctx := context.Background()
			collection := database.MongoDb.Collection(os.Getenv("GAME_VALORANT_SKINS_COLLECTION_NAME"))

			var userData models.UserSkinGame
			err = collection.FindOne(ctx, bson.D{{"_id", i.Member.User.ID}}).Decode(&userData)
			if err != nil {
				if errors.Is(err, mongo.ErrNoDocuments) {
					newUser := models.UserSkinGame{UserId: i.Member.User.ID, RollRemaining: 5}
					_, err = collection.InsertOne(ctx, &newUser)
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

			if userData.RollRemaining == 0 {
				err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Embeds: embed.New().
							SetTitle(i18n.Get(discordgo.French, "game_valorant_skins.roll_no_more_roll_title")).
							SetDescription(i18n.Get(discordgo.French, "game_valorant_skins.roll_no_more_roll_description", i18n.Vars{
								"timeRemaining": getResetRollTimeRemaining(),
							})).
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

			randomInt := rand.Intn(len(skins.Data))
			_, err = s.ChannelMessageSendEmbeds(i.ChannelID, embed.New().
				SetTitle(i18n.Get(discordgo.French, "game_valorant_skins.roll_response_title", i18n.Vars{
					"rollAuthorUsername": i.Member.User.GlobalName,
				})).
				SetColor(embed.VALOROUS).
				SetCurrentTimestamp().
				SetDefaultFooter().
				AddInlinedField("ℹ️ Skin", skins.Data[randomInt].DisplayName).
				AddInlinedField("💎 Rareté", "❔").
				SetThumbnail("https://zupimages.net/up/24/22/b5js.png").
				SetImage(skins.Data[randomInt].DisplayIcon).
				ToMessageEmbeds())
			if err != nil {
				log.Logger.Error(err)
				return
			}

			userData.RollRemaining = userData.RollRemaining - 1

			for _, skin := range userData.SkinsCollectedName {
				if skin.DisplayName == skins.Data[randomInt].DisplayName {
					err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
						Type: discordgo.InteractionResponseChannelMessageWithSource,
						Data: &discordgo.InteractionResponseData{
							Embeds: embed.New().
								SetTitle(i18n.Get(discordgo.French, "game_valorant_skins.roll_private_response_title")).
								SetDescription(i18n.Get(discordgo.French, "game_valorant_skins.roll_private_response_already_have_description", i18n.Vars{
									"rollRemaining": userData.RollRemaining,
									"timeRemaining": getResetRollTimeRemaining(),
								})).
								SetColor(embed.VALOROUS).
								SetThumbnail("https://zupimages.net/up/24/22/0jcs.png").
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

					err = saveUserData(ctx, userData, collection)
					if err != nil {
						log.Logger.Error(err)
						return
					}
					return
				}
			}

			userData.SkinsCollectedName = append(userData.SkinsCollectedName, skins.Data[randomInt])

			err = saveUserData(ctx, userData, collection)
			if err != nil {
				log.Logger.Error(err)
				return
			}

			err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: embed.New().
						SetTitle(i18n.Get(discordgo.French, "game_valorant_skins.roll_private_response_title")).
						SetDescription(i18n.Get(discordgo.French, "game_valorant_skins.roll_private_response_new_skin_description", i18n.Vars{
							"rollRemaining": userData.RollRemaining,
							"timeRemaining": getResetRollTimeRemaining(),
						})).
						SetColor(embed.VALOROUS).
						SetCurrentTimestamp().
						SetThumbnail("https://zupimages.net/up/24/22/fgqy.png").
						SetDefaultFooter().
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

func saveUserData(ctx context.Context, user models.UserSkinGame, collection *mongo.Collection) error {
	_, err := collection.UpdateOne(ctx, bson.D{{"_id", user.UserId}}, bson.M{"$set": &user})
	if err != nil {
		return err
	}
	return nil
}

func getResetRollTimeRemaining() string {
	now := time.Now()
	midnight := time.Date(now.Year(), now.Month(), now.Day(), 24, 0, 0, 0, now.Location())
	durationUntilMidnight := midnight.Sub(now)
	return fmt.Sprintf("%02dh %02dm %02ds", int(durationUntilMidnight.Hours()), int(durationUntilMidnight.Minutes())%60, int(durationUntilMidnight.Seconds())%60)
}
