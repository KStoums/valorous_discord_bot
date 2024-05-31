package models

type WeaponSkin struct {
	UUID            string   `json:"uuid"`
	DisplayName     string   `json:"displayName"`
	ThemeUUID       string   `json:"themeUuid"`
	ContentTierUUID string   `json:"contentTierUuid"`
	DisplayIcon     string   `json:"displayIcon"`
	Wallpaper       string   `json:"wallpaper"`
	AssetPath       string   `json:"assetPath"`
	Chromas         []chroma `json:"chromas"`
	Levels          []level  `json:"levels"`
}

type chroma struct {
	UUID          string `json:"uuid"`
	DisplayName   string `json:"displayName"`
	DisplayIcon   string `json:"displayIcon"`
	FullRender    string `json:"fullRender"`
	Swatch        string `json:"swatch"`
	StreamedVideo string `json:"streamedVideo"`
	AssetPath     string `json:"assetPath"`
}

type level struct {
	UUID          string `json:"uuid"`
	DisplayName   string `json:"displayName"`
	LevelItem     string `json:"levelItem"`
	DisplayIcon   string `json:"displayIcon"`
	StreamedVideo string `json:"streamedVideo"`
	AssetPath     string `json:"assetPath"`
}
