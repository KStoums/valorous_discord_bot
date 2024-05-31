package task

import (
	"context"
	"github.com/bwmarrin/discordgo"
	"github.com/goroutine/template/database"
	"github.com/goroutine/template/log"
	"github.com/goroutine/template/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

type AddRolls struct {
	s *discordgo.Session
}

func NewAddRollsTask(s *discordgo.Session) *AddRolls {
	return &AddRolls{s: s}
}

func (a *AddRolls) Run() {
	ctx := context.Background()
	collection := database.MongoDb.Collection(os.Getenv("GAME_VALORANT_SKINS_COLLECTION_NAME"))
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Logger.Error(err)
		return
	}
	defer cur.Close(ctx)

	var usersData []models.UserSkinGame
	err = cur.All(ctx, &usersData)
	if err != nil {
		log.Logger.Error(err)
		return
	}

	for _, user := range usersData {
		user.RollRemaining = 5
		err = a.saveUserData(ctx, user, collection)
		if err != nil {
			log.Logger.Error(err)
			return
		}
	}

	log.Logger.Info("Rolls redefined to 5 to alls users")
}

func (a *AddRolls) CronString() string {
	return "0 0 * * *"
}

func (a *AddRolls) Name() string {
	return "AddRolls"
}
func (a *AddRolls) RunOnStart() bool {
	return false
}

func (a *AddRolls) saveUserData(ctx context.Context, user models.UserSkinGame, collection *mongo.Collection) error {
	_, err := collection.UpdateOne(ctx, bson.D{{"_id", user.UserId}}, bson.M{"$set": &user})
	if err != nil {
		return err
	}
	return nil
}
