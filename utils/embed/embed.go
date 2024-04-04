package embed

import (
	"github.com/bwmarrin/discordgo"
	"github.com/goccy/go-json"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type Embed struct {
	discordgo.MessageEmbed
}

const (
	PURPLE_DARK   = "#6a006a"
	PURPLE_MEDIUM = "#a958a5"
	PURPLE_LIGHT  = "#c481fb"
	ORANGE        = "#ffa500"
	GOLD          = "#daa520"
	RED_DARK      = "#8e2430"
	RED_LIGHT     = "#f94343"
	BLUE_DARK     = "#3b5998"
	CYAN          = "#5780cd"
	BLUE_LIGHT    = "#ace9e7"
	AQUA          = "#33a1ee"
	PINK          = "#ff9dbb"
	GREEN_DARK    = "#2ac075"
	GREEN_LIGHT   = "#a1ee33"
	WHITE         = "#f9f9f6"
	CREAM         = "#ffdab9"
	ELORION       = "#ff9900"
	VALOROUS      = "#652481"
)

func New() *Embed {
	return &Embed{}
}

func (e *Embed) AddField(name, value string) *Embed {
	e.Fields = append(e.Fields, &discordgo.MessageEmbedField{
		Name:  name,
		Value: value,
	})
	return e
}

func (e *Embed) AddInlinedField(name, value string) *Embed {
	e.Fields = append(e.Fields, &discordgo.MessageEmbedField{
		Name:   name,
		Value:  value,
		Inline: true,
	})
	return e
}

func (e *Embed) Equal(embed *Embed) bool {
	return reflect.DeepEqual(e, embed)
}

func (e *Embed) SetAuthor(name, iconUrl, url string) *Embed {
	e.Author = &discordgo.MessageEmbedAuthor{
		URL:     url,
		Name:    name,
		IconURL: iconUrl,
	}
	return e
}

func (e *Embed) SetColor(color string) *Embed {
	parsedColor, err := strconv.ParseInt(strings.Replace(color, "#", "", 1), 16, 64)
	if err == nil {
		e.Color = int(parsedColor)
	}
	return e
}

func (e *Embed) SetDescription(description string) *Embed {
	e.Description = description
	return e
}

func (e *Embed) SetFooter(text, iconUrl string) *Embed {
	e.Footer = &discordgo.MessageEmbedFooter{
		Text:    text,
		IconURL: iconUrl,
	}
	return e
}

func (e *Embed) SetDefaultFooter() *Embed {
	e.Footer = &discordgo.MessageEmbedFooter{
		Text: "Valorous",
	}
	return e
}

func (e *Embed) SetImage(url string) *Embed {
	e.Image = &discordgo.MessageEmbedImage{URL: url}
	return e
}

func (e *Embed) SetThumbnail(url string) *Embed {
	e.Thumbnail = &discordgo.MessageEmbedThumbnail{URL: url}
	return e
}

func (e *Embed) SetTimestamp(t time.Time) *Embed {
	e.Timestamp = t.UTC().Format("2006-01-02T15:04:05-0700")
	return e
}

func (e *Embed) SetCurrentTimestamp() *Embed {
	e.Timestamp = time.Now().Format("2006-01-02T15:04:05-0700")
	return e
}

func (e *Embed) SetTitle(title string) *Embed {
	e.Title = title
	return e
}

func (e *Embed) SetUrl(url string) *Embed {
	e.URL = url
	return e
}

func (e *Embed) ToJson() (string, error) {
	jsonResult, err := json.Marshal(e)
	if err != nil {
		return "", err
	}

	return string(jsonResult), nil
}

func (e *Embed) SetVideo(url string) *Embed {
	e.Video = &discordgo.MessageEmbedVideo{
		URL:    url,
		Width:  100,
		Height: 50,
	}
	return e
}

func (e *Embed) ToMessageEmbed() *discordgo.MessageEmbed {
	return &e.MessageEmbed
}

func (e *Embed) ToMessageEmbeds() []*discordgo.MessageEmbed {
	return []*discordgo.MessageEmbed{&e.MessageEmbed}
}
