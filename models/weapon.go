package models

type Weapons struct {
	UUID            string      `json:"uuid"`
	DisplayName     string      `json:"displayName"`
	Category        string      `json:"category"`
	DefaultSkinUUID string      `json:"defaultSkinUuid"`
	DisplayIcon     string      `json:"displayIcon"`
	KillStreamIcon  string      `json:"killStreamIcon"`
	AssetPath       string      `json:"assetPath"`
	WeaponStats     WeaponStats `json:"weaponStats"`
	ShopData        ShopData    `json:"shopData"`
	Skins           []Skin      `json:"skins"`
}

type WeaponStats struct {
	FireRate            float64 `json:"fireRate"`
	MagazineSize        int     `json:"magazineSize"`
	RunSpeedMultiplier  float64 `json:"runSpeedMultiplier"`
	EquipTimeSeconds    float64 `json:"equipTimeSeconds"`
	ReloadTimeSeconds   float64 `json:"reloadTimeSeconds"`
	FirstBulletAccuracy float64 `json:"firstBulletAccuracy"`
	ShotgunPelletCount  int     `json:"shotgunPelletCount"`
	WallPenetration     string  `json:"wallPenetration"`
	Feature             string  `json:"feature"`
	FireMode            *string `json:"fireMode"`
	AltFireType         string  `json:"altFireType"`
	AdsStats            struct {
		ZoomMultiplier      float64 `json:"zoomMultiplier"`
		FireRate            float64 `json:"fireRate"`
		RunSpeedMultiplier  float64 `json:"runSpeedMultiplier"`
		BurstCount          int     `json:"burstCount"`
		FirstBulletAccuracy float64 `json:"firstBulletAccuracy"`
	} `json:"adsStats"`
	AltShotgunStats *interface{} `json:"altShotgunStats"`
	AirBurstStats   *interface{} `json:"airBurstStats"`
	DamageRanges    []struct {
		RangeStartMeters float64 `json:"rangeStartMeters"`
		RangeEndMeters   float64 `json:"rangeEndMeters"`
		HeadDamage       float64 `json:"headDamage"`
		BodyDamage       float64 `json:"bodyDamage"`
		LegDamage        float64 `json:"legDamage"`
	} `json:"damageRanges"`
}

type GridPosition struct {
	Row    int `json:"row"`
	Column int `json:"column"`
}

type ShopData struct {
	Cost              int          `json:"cost"`
	Category          string       `json:"category"`
	ShopOrderPriority int          `json:"shopOrderPriority"`
	CategoryText      string       `json:"categoryText"`
	GridPosition      GridPosition `json:"gridPosition"`
	CanBeTrashed      bool         `json:"canBeTrashed"`
	Image             *string      `json:"image"`
	NewImage          string       `json:"newImage"`
	NewImage2         *string      `json:"newImage2"`
	AssetPath         string       `json:"assetPath"`
}

type Skin struct {
	UUID            string   `json:"uuid"`
	DisplayName     string   `json:"displayName"`
	ThemeUUID       string   `json:"themeUuid"`
	ContentTierUUID string   `json:"contentTierUuid"`
	DisplayIcon     string   `json:"displayIcon"`
	Wallpaper       *string  `json:"wallpaper"`
	AssetPath       string   `json:"assetPath"`
	Chromas         []Chroma `json:"chromas"`
	Levels          []Level  `json:"levels"`
}

type Chroma struct {
	UUID          string  `json:"uuid"`
	DisplayName   string  `json:"displayName"`
	DisplayIcon   *string `json:"displayIcon"`
	FullRender    string  `json:"fullRender"`
	Swatch        *string `json:"swatch"`
	StreamedVideo *string `json:"streamedVideo"`
	AssetPath     string  `json:"assetPath"`
}

type Level struct {
	UUID          string  `json:"uuid"`
	DisplayName   string  `json:"displayName"`
	LevelItem     *string `json:"levelItem"`
	DisplayIcon   string  `json:"displayIcon"`
	StreamedVideo *string `json:"streamedVideo"`
	AssetPath     string  `json:"assetPath"`
}
