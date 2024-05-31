package config

// Config Define config here with substructs for each config category
type Config struct {
	GuildId  string   `json:"guildId"`
	Channels channels `json:"channels"`
	Roles    roles    `json:"roles"`
}

type channels struct {
	MemberCount            string `json:"memberCount"`
	WelcomeChannel         string `json:"welcomeChannel"`
	RulesChannel           string `json:"rulesChannel"`
	AutoVoiceChannel       string `json:"autoVoiceChannel"`
	AutoVoiceChannelTeam   string `json:"autoVoiceChannelTeam"`
	AutoVoiceChannelAdmin  string `json:"autoVoiceChannelAdmin"`
	TeamCategory           string `json:"teamCategory"`
	AutoVoiceCategory      string `json:"autoVoiceCategory"`
	AutoRoleRankedChannel  string `json:"autoRoleRankedChannel"`
	LogChannel             string `json:"logChannel"`
	SuggestionChannel      string `json:"suggestionChannel"`
	ClipChannel            string `json:"clipChannel"`
	PublicAnnounceChannel  string `json:"publicAnnounceChannel"`
	TeamAnnounceChannel    string `json:"teamAnnounceChannel"`
	TicketArchivedCategory string `json:"ticketArchivedCategory"`
	TicketOpenCategory     string `json:"ticketOpenCategory"`
	TicketChannel          string `json:"ticketChannel"`
	AdminCategory          string `json:"adminCategory"`
	BotCommand             string `json:"botCommand"`
	BotCommandAdmin        string `json:"botCommandAdmin"`
	BotChangelog           string `json:"botChangelog"`
}

type roles struct {
	AdministrationRole string `json:"administrationRole"`
	MemberRole         string `json:"memberRole"`
	AcceptedRuleRole   string `json:"acceptedRuleRole"`
	UnrankedRole       string `json:"unrankedRole"`
	IronRole           string `json:"ironRole"`
	BronzeRole         string `json:"bronzeRole"`
	SilverRole         string `json:"silverRole"`
	GoldRole           string `json:"goldRole"`
	PlatinumRole       string `json:"platinumRole"`
	DiamondRole        string `json:"diamondRole"`
	AscendantRole      string `json:"ascendantRole"`
	ImmortalRole       string `json:"immortalRole"`
	RadiantRole        string `json:"radiantRole"`
	RankSeparatorRole  string `json:"rankSeparatorRole"`
	MutedRole          string `json:"mutedRole"`
}
