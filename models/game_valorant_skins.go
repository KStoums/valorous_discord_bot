package models

type UserSkinGame struct {
	UserId             string       `json:"userId" gorm:"primary_key;not null;unique" bson:"_id"`
	SkinsCollectedName []WeaponSkin `json:"skinsCollectedName" bson:"skinsCollectedName"`
	RollRemaining      int          `json:"rollRemaining" bson:"rollRemaining"`
}
