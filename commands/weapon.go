package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/goccy/go-json"
	"github.com/goroutine/template/log"
	"github.com/goroutine/template/models"
	"github.com/goroutine/template/utils/embed"
	i18n "github.com/kaysoro/discordgo-i18n"
	"net/http"
	"strconv"
	"strings"
)

const valorantApiWeaponUrl = "https://valorant-api.com/v1/weapons"

func WeaponCommand() SlashCommand {
	return SlashCommand{
		Name:        "weapon",
		Description: i18n.Get(discordgo.French, "weapon.command_description"),
		Enabled:     true,
		ArgsFunc: ArgsFromStructs(
			SlashCommandArg{
				Name:        "weapon_name",
				Description: i18n.Get(discordgo.French, "weapon.weapon_args"),
				Type:        discordgo.ApplicationCommandOptionString,
				Required:    true,
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "Vandal",
						Value: "vandal",
					},
					{
						Name:  "Phantom",
						Value: "phantom",
					},
					{
						Name:  "Guardian",
						Value: "guardian",
					},
					{
						Name:  "Bulldog",
						Value: "bulldog",
					},
					{
						Name:  "Odin",
						Value: "odin",
					},
					{
						Name:  "Ares",
						Value: "ares",
					},
					{
						Name:  "Operator",
						Value: "operator",
					},
					{
						Name:  "Outlaw",
						Value: "outlaw",
					},
					{
						Name:  "Marshal",
						Value: "marshal",
					},
					{
						Name:  "Judge",
						Value: "judge",
					},
					{
						Name:  "Bucky",
						Value: "bucky",
					},
					{
						Name:  "Spectre",
						Value: "spectre",
					},
					{
						Name:  "Stinger",
						Value: "stinger",
					},
					{
						Name:  "Sherif",
						Value: "sherif",
					},
					{
						Name:  "Ghost",
						Value: "ghost",
					},
					{
						Name:  "Frenzy",
						Value: "frenzy",
					},
					{
						Name:  "Shorty",
						Value: "shorty",
					},
					{
						Name:  "Classic",
						Value: "classic",
					},
					{
						Name:  "Couteau",
						Value: "knife",
					},
				},
			},
		),
		Handler: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			weaponName := i.ApplicationCommandData().Options[0].StringValue()

			resp, err := http.Get(valorantApiWeaponUrl)
			if err != nil {
				log.Logger.Error(err)
				return
			}

			var responseBody struct {
				Status int32            `json:"status"`
				Data   []models.Weapons `json:"data"`
			}
			err = json.NewDecoder(resp.Body).Decode(&responseBody)
			if err != nil {
				log.Logger.Error(err)
				return
			}

			var weapon models.Weapons
			for _, w := range responseBody.Data {
				if strings.ToLower(w.DisplayName) == weaponName {
					weapon = w
					break
				}
			}

			penetrationSplit := strings.SplitAfter(weapon.WeaponStats.WallPenetration, "::")

			finalEmbeds := embed.New().
				SetTitle(i18n.Get(discordgo.French, "weapon.weapon_found_title", i18n.Vars{
					"weaponName": weapon.DisplayName,
				})).
				SetColor(embed.VALOROUS).
				SetCurrentTimestamp().
				SetDefaultFooter().
				SetImage(weapon.DisplayIcon).
				SetThumbnail("https://zupimages.net/up/24/16/yte5.png").
				AddInlinedField("💰 Prix", strconv.Itoa(weapon.ShopData.Cost)+"$").
				AddInlinedField("📈 Cadence", strconv.Itoa(int(weapon.WeaponStats.FireRate))+" balles par seconde(s)").
				AddInlinedField("👝 Chargeur", strconv.Itoa(weapon.WeaponStats.MagazineSize)+" balles").
				AddInlinedField("⏱️ Chargement", strconv.Itoa(int(weapon.WeaponStats.ReloadTimeSeconds))+" seconde(s)").
				AddInlinedField("⏱️ Sortie", strconv.Itoa(int(weapon.WeaponStats.EquipTimeSeconds))+" seconde(s)").
				AddInlinedField("🧱 Pénétration", penetrationSplit[1]).
				ToMessageEmbeds()

			for _, damageRange := range weapon.WeaponStats.DamageRanges {
				finalEmbeds[0].Fields = append(finalEmbeds[0].Fields, &discordgo.MessageEmbedField{
					Name:   fmt.Sprintf("🤕 Tête \n(%v-%vm)", damageRange.RangeStartMeters, damageRange.RangeEndMeters),
					Value:  strconv.Itoa(int(damageRange.HeadDamage)),
					Inline: true,
				})

				finalEmbeds[0].Fields = append(finalEmbeds[0].Fields, &discordgo.MessageEmbedField{
					Name:   fmt.Sprintf("👃 Corp \n(%v-%vm)", damageRange.RangeStartMeters, damageRange.RangeEndMeters),
					Value:  strconv.Itoa(int(damageRange.BodyDamage)),
					Inline: true,
				})

				finalEmbeds[0].Fields = append(finalEmbeds[0].Fields, &discordgo.MessageEmbedField{
					Name:   fmt.Sprintf("🦶 Jambes \n(%v-%vm)", damageRange.RangeStartMeters, damageRange.RangeEndMeters),
					Value:  strconv.Itoa(int(damageRange.LegDamage)),
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
