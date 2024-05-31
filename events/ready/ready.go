package ready

import (
	"github.com/bwmarrin/discordgo"
	"github.com/goroutine/template/commands"
	"github.com/goroutine/template/features/auto_role"
	"github.com/goroutine/template/features/rule"
	"github.com/goroutine/template/features/ticket"
	"github.com/goroutine/template/log"
	"github.com/goroutine/template/task"
)

func ReadyEvent(s *discordgo.Session, r *discordgo.Ready) {
	log.Logger.Info("Template Bot is now running. Press CTRL+C to exit.")

	rule.CreateRuleEmbed(s)
	auto_role.CreateAutoRoleRankedEmbed(s)
	ticket.CreateTicketEmbed(s)

	taskManager := task.NewTaskManager()
	taskManager.AddTasks(task.NewGuildMemberCountTask(s), task.NewBotPresence(s), task.NewAddRollsTask(s))
	taskManager.RunTasks()

	commands.RegisterCommands(s, r.User)
}
