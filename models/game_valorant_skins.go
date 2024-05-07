package models

type GameValorantSkins struct {
	UserId             string   `json:"userId" gorm:"primary_key;not null;unique" bson:"userId"`
	SkinsCollectedName []string `json:"skinsCollectedName" bson:"skinsCollectedName"`
}
