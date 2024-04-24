package valorant

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/goccy/go-json"
	"github.com/goroutine/template/commands"
	"github.com/goroutine/template/log"
	"github.com/goroutine/template/models"
	"github.com/goroutine/template/utils/embed"
	i18n "github.com/kaysoro/discordgo-i18n"
	"net/http"
	"strconv"
	"strings"
)

const valorantApiWeaponUrl = "https://valorant-api.com/v1/weapons?language=fr-FR"

var weaponsCommandArgs []*discordgo.ApplicationCommandOptionChoice

type responseBodyWeapon struct {
	Status int32           `json:"status"`
	Data   []models.Weapon `json:"data"`
}

func init() {
	weapons, err := requestWeaponsApi()
	if err != nil {
		panic(err)
	}

	for _, w := range weapons.Data {
		weaponsCommandArgs = append(weaponsCommandArgs, &discordgo.ApplicationCommandOptionChoice{
			Name:  w.DisplayName,
			Value: w.DisplayName,
		})
	}
}

func WeaponCommand() commands.SlashCommand {
	return commands.SlashCommand{
		Name:        "weapon",
		Description: i18n.Get(discordgo.French, "weapon.command_description"),
		Enabled:     true,
		ArgsFunc: commands.ArgsFromStructs(
			commands.SlashCommandArg{
				Name:        "weapon_name",
				Description: i18n.Get(discordgo.French, "weapon.weapon_args"),
				Type:        discordgo.ApplicationCommandOptionString,
				Required:    true,
				Choices:     weaponsCommandArgs,
			},
		),
		Handler: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			weaponName := i.ApplicationCommandData().Options[0].StringValue()

			weapons, err := requestWeaponsApi()
			if err != nil {
				log.Logger.Error(err)
				return
			}

			var weapon models.Weapon
			for _, w := range weapons.Data {
				if w.DisplayName == weaponName {
					weapon = w
					break
				}
			}

			finalEmbeds := embed.New().
				SetTitle(i18n.Get(discordgo.French, "weapon.weapon_title", i18n.Vars{
					"weaponName": weaponName,
				})).
				SetColor(embed.VALOROUS).
				SetCurrentTimestamp().
				SetDefaultFooter().
				SetImage(weapon.DisplayIcon).
				SetThumbnail("https://zupimages.net/up/24/16/yte5.png").
				AddInlinedField("💰 Prix", strconv.Itoa(weapon.ShopData.Cost)+"$").
				AddInlinedField("💥 Cadence", strconv.Itoa(int(weapon.WeaponStats.FireRate))+" balle(s) par seconde(s)").
				AddInlinedField("👝 Chargeur", strconv.Itoa(weapon.WeaponStats.MagazineSize)+" balle(s)").
				AddInlinedField("⏱️ Chargement", strconv.Itoa(int(weapon.WeaponStats.ReloadTimeSeconds))+" seconde(s)").
				AddInlinedField("⏱️ Sortie", strconv.Itoa(int(weapon.WeaponStats.EquipTimeSeconds))+" seconde(s)").
				ToMessageEmbeds()

			if weapon.WeaponStats.WallPenetration != "" {
				finalEmbeds[0].Fields = append(finalEmbeds[0].Fields, &discordgo.MessageEmbedField{
					Name:   "🧱 Pénétration",
					Value:  strings.Split(weapon.WeaponStats.WallPenetration, "::")[1],
					Inline: true,
				})
			}

			for _, damageRange := range weapon.WeaponStats.DamageRanges {
				finalEmbeds[0].Fields = append(finalEmbeds[0].Fields, &discordgo.MessageEmbedField{
					Name:   fmt.Sprintf("🤕 Tête \n(%v-%vm)", damageRange.RangeStartMeters, damageRange.RangeEndMeters),
					Value:  strconv.Itoa(int(damageRange.HeadDamage)),
					Inline: true,
				})

				finalEmbeds[0].Fields = append(finalEmbeds[0].Fields, &discordgo.MessageEmbedField{
					Name:   fmt.Sprintf("🫃🏻 Corp \n(%v-%vm)", damageRange.RangeStartMeters, damageRange.RangeEndMeters),
					Value:  strconv.Itoa(int(damageRange.BodyDamage)),
					Inline: true,
				})

				finalEmbeds[0].Fields = append(finalEmbeds[0].Fields, &discordgo.MessageEmbedField{
					Name:   fmt.Sprintf("🦵 Jambes \n(%v-%vm)", damageRange.RangeStartMeters, damageRange.RangeEndMeters),
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

func requestWeaponsApi() (responseBodyWeapon, error) {
	resp, err := http.Get(valorantApiWeaponUrl)
	if err != nil {
		log.Logger.Error(err)
		return responseBodyWeapon{}, err
	}
	defer resp.Body.Close()

	var response responseBodyWeapon
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Logger.Error(err)
		return responseBodyWeapon{}, err
	}

	return response, nil
}
