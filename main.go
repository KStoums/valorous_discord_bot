package main

import (
	"context"
	"github.com/bwmarrin/discordgo"
	"github.com/goroutine/template/commands"
	"github.com/goroutine/template/commands/game/valorant_skins"
	"github.com/goroutine/template/commands/moderation"
	"github.com/goroutine/template/commands/valorant"
	"github.com/goroutine/template/config"
	"github.com/goroutine/template/events"
	"github.com/goroutine/template/events/ready"
	"github.com/goroutine/template/log"
	"github.com/goroutine/template/utils"
	"github.com/goroutine/template/utils/environnement"
	_ "github.com/joho/godotenv/autoload" // Load .env file
	i18n "github.com/kaysoro/discordgo-i18n"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	ctx := context.Background()

	environnement.CheckEnvs() // Check if all envs are set
	log.Logger.Info("Currently running commit: " + utils.GetCommit())

	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		panic(err)
	}
	defer mongoClient.Disconnect(context.Background())

	err = mongoClient.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}

	database := mongoClient.Database(strings.ToLower(os.Getenv("PROJECT_NAME")))

	err = i18n.LoadBundle(discordgo.French, "./locales/fr.json")
	if err != nil {
		log.Logger.Fatal(err)
	}

	config.LoadConfig()

	discord, err := discordgo.New("Bot " + environnement.GetToken())
	if err != nil {
		log.Logger.Fatal(err)
	}

	discord.Identify.Intents = discordgo.IntentsAll

	commands.AddCommands(moderation.ClearCommand(), moderation.MuteCommand(), moderation.UnmuteCommand(),
		valorant.WeaponCommand(), valorant.MapCommand(), valorant.AgentCommand(), valorant_skins.SkinDexCommand(database),
		valorant_skins.RollCommand(database))

	discord.AddHandlerOnce(ready.ReadyEvent)
	addHandlers(discord, events.InteractionCreateEvent, events.MemberJoinEvent, events.VoiceStateUpdateEvent,
		events.ReactionAddEvent, events.ReactionRemoveEvent, events.MessageCreateEvent)

	if err = discord.Open(); err != nil {
		log.Logger.Fatal(err)
	}
	defer discord.Close()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

func addHandlers(s *discordgo.Session, handlers ...interface{}) {
	for _, handler := range handlers {
		s.AddHandler(handler)
	}
}
