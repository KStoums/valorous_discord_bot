package models

type Agent struct {
	UUID                      string          `json:"uuid"`
	DisplayName               string          `json:"displayName"`
	Description               string          `json:"description"`
	DeveloperName             string          `json:"developerName"`
	CharacterTags             []string        `json:"characterTags"`
	DisplayIcon               string          `json:"displayIcon"`
	DisplayIconSmall          string          `json:"displayIconSmall"`
	BustPortrait              string          `json:"bustPortrait"`
	FullPortrait              string          `json:"fullPortrait"`
	FullPortraitV2            string          `json:"fullPortraitV2"`
	KillFeedPortrait          string          `json:"killfeedPortrait"`
	Background                string          `json:"background"`
	BackgroundGradientColors  []string        `json:"backgroundGradientColors"`
	AssetPath                 string          `json:"assetPath"`
	IsFullPortraitRightFacing bool            `json:"isFullPortraitRightFacing"`
	IsPlayableCharacter       bool            `json:"isPlayableCharacter"`
	IsAvailableForTest        bool            `json:"isAvailableForTest"`
	IsBaseContent             bool            `json:"isBaseContent"`
	Role                      Role            `json:"role"`
	RecruitmentData           RecruitmentData `json:"recruitmentData"`
	Abilities                 []Ability       `json:"abilities"`
}

type Role struct {
	UUID        string `json:"uuid"`
	DisplayName string `json:"displayName"`
	Description string `json:"description"`
	DisplayIcon string `json:"displayIcon"`
	AssetPath   string `json:"assetPath"`
}

type VoiceLine struct {
	MinDuration float32 `json:"minDuration"`
	MaxDuration float32 `json:"maxDuration"`
}

type Media struct {
	ID    int32  `json:"id"`
	Wwise string `json:"wwise"`
	Wave  string `json:"wave"`
}

type Ability struct {
	Slot        string    `json:"slot"`
	DisplayName string    `json:"displayName"`
	Description string    `json:"description"`
	DisplayIcon string    `json:"displayIcon"`
	VoiceLine   VoiceLine `json:"voiceLine"`
	MediaList   []Media   `json:"mediaList"`
}

type RecruitmentData struct {
	CounterID              string `json:"counterId"`
	MilestoneID            string `json:"milestoneId"`
	MilestoneThreshold     int32  `json:"milestoneThreshold"`
	UseLevelVpCostOverride bool   `json:"useLevelVpCostOverride"`
	LevelVpCostOverride    int32  `json:"levelVpCostOverride"`
	StartDate              string `json:"startDate"`
	EndDate                string `json:"endDate"`
}
