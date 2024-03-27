package ready

import (
	"github.com/bwmarrin/discordgo"
	"github.com/goroutine/template/commands"
	"github.com/goroutine/template/features/rules"
	"github.com/goroutine/template/log"
	"github.com/goroutine/template/task"
)

func ReadyEvent(s *discordgo.Session, r *discordgo.Ready) {
	log.Logger.Info("Template Bot is now running. Press CTRL+C to exit.")

	rules.CreateRulesEmbed(s)

	taskManager := task.NewTaskManager()
	taskManager.AddTasks(task.NewGuildMemberCountTask(s))
	taskManager.RunTasks()

	commands.RegisterCommands(s, r.User)
}
