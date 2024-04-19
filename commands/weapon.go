package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/goccy/go-json"
	"github.com/goroutine/template/log"
	"github.com/goroutine/template/models"
	"github.com/goroutine/template/utils/embed"
	i18n "github.com/kaysoro/discordgo-i18n"
	"io"
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
				err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Embeds: embed.New().
							SetTitle(i18n.Get(discordgo.French, "weapon.request_valorant_api_error_title")).
							SetDescription(i18n.Get(discordgo.French, "weapon.request_valorant_api_error_description")).
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

			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Logger.Error(err)
				return
			}

			var requestResponse []struct {
				Status int            `json:"status"`
				Data   models.Weapons `json:"data"`
			}
			err = json.Unmarshal(bodyBytes, &requestResponse)
			if err != nil {
				log.Logger.Error(err)
				return
			}

			var weapon models.Weapons
			for _, w := range requestResponse {
				if strings.ToLower(w.Data.DisplayName) == weaponName {
					weapon = w.Data
					break
				}
			}

			err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: embed.New().
						SetTitle(i18n.Get(discordgo.French, "weapon.weapon_found_title")).
						SetColor(embed.VALOROUS).
						SetCurrentTimestamp().
						SetDefaultFooter().
						SetThumbnail(weapon.DisplayIcon).
						AddInlinedField("💰 Prix", strconv.Itoa(weapon.ShopData.Cost)).
						AddInlinedField("📈 Cadence de tir", strconv.Itoa(int(weapon.WeaponStats.FireRate))).
						AddInlinedField("👝 Taille du chargeur", strconv.Itoa(weapon.WeaponStats.MagazineSize)).
						AddInlinedField("⏱️ Temps de rechargement", strconv.Itoa(int(weapon.WeaponStats.ReloadTimeSeconds))).
						AddInlinedField("⏱️ Temps à l'équipement", strconv.Itoa(int(weapon.WeaponStats.EquipTimeSeconds))).
						AddInlinedField("🧱 Pénétration des balles", weapon.WeaponStats.WallPenetration).
						AddInlinedField("🤕 Dégâts dans la tête (0-30m)", strconv.Itoa(int(weapon.WeaponStats.DamageRanges[0].HeadDamage))).
						AddInlinedField("👃 Dégâts dans le corp (0-30m)", strconv.Itoa(int(weapon.WeaponStats.DamageRanges[0].BodyDamage))).
						AddInlinedField("🦶 Dégâts dans les jambes (0-30m)", strconv.Itoa(int(weapon.WeaponStats.DamageRanges[0].LegDamage))).
						AddInlinedField("🤕 Dégâts dans la tête (30-50m)", strconv.Itoa(int(weapon.WeaponStats.DamageRanges[1].LegDamage))).
						AddInlinedField("👃 Dégâts dans le corp (30-50m)", strconv.Itoa(int(weapon.WeaponStats.DamageRanges[1].BodyDamage))).
						AddInlinedField("🦶 Dégâts dans les jambes (30-50m)", strconv.Itoa(int(weapon.WeaponStats.DamageRanges[1].LegDamage))).
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
