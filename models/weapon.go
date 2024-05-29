package models

type Weapon struct {
	UUID        string      `json:"uuid"`
	DisplayName string      `json:"displayName"`
	Category    string      `json:"category"`
	DisplayIcon string      `json:"displayIcon"`
	WeaponStats weaponStats `json:"weaponStats"`
	ShopData    shopData    `json:"shopData"`
}

type weaponStats struct {
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

type gridPosition struct {
	Row    int `json:"row"`
	Column int `json:"column"`
}

type shopData struct {
	Cost              int          `json:"cost"`
	Category          string       `json:"category"`
	ShopOrderPriority int          `json:"shopOrderPriority"`
	CategoryText      string       `json:"categoryText"`
	GridPosition      gridPosition `json:"gridPosition"`
	CanBeTrashed      bool         `json:"canBeTrashed"`
	Image             *string      `json:"image"`
	NewImage          string       `json:"newImage"`
	NewImage2         *string      `json:"newImage2"`
	AssetPath         string       `json:"assetPath"`
}
