package models

type Map struct {
	UUID                  string    `json:"uuid"`
	DisplayName           string    `json:"displayName"`
	NarrativeDescription  string    `json:"narrativeDescription"`
	TacticalDescription   string    `json:"tacticalDescription"`
	Coordinates           string    `json:"coordinates"`
	DisplayIcon           string    `json:"displayIcon"`
	ListViewIcon          string    `json:"listViewIcon"`
	ListViewIconTall      string    `json:"listViewIconTall"`
	Splash                string    `json:"splash"`
	StylizedBackgroundImg string    `json:"stylizedBackgroundImage"`
	PremierBackgroundImg  string    `json:"premierBackgroundImage"`
	AssetPath             string    `json:"assetPath"`
	MapURL                string    `json:"mapUrl"`
	XMultiplier           float32   `json:"xMultiplier"`
	YMultiplier           float32   `json:"yMultiplier"`
	XScalarToAdd          float32   `json:"xScalarToAdd"`
	YScalarToAdd          float32   `json:"yScalarToAdd"`
	Callouts              []callout `json:"callouts"`
}

type localizedString struct {
	Default string `json:"default"`
	// Add other localized fields as needed
}

type callout struct {
	RegionName      string `json:"regionName"`
	SuperRegionName string `json:"superRegionName"`
	Location        struct {
		X float32 `json:"x"`
		Y float32 `json:"y"`
	} `json:"location"`
}
