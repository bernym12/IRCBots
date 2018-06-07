package dad

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/adammohammed/groupmebot"
)

// Run runs
func (gb *GroupMeBot) Run(groupmeConfigFileName, botConfigFileName string) {
	rand.Seed(time.Now().Unix())

	lg := groupmebot.CSVLogger{LogFile: "bot.csv"}
	stdout := groupmebot.StdOutLogger{}
	combinedLogger := groupmebot.CompositeLogger{Loggers: []groupmebot.Logger{lg, stdout}}

	bot := groupmebot.GroupMeBot{Logger: combinedLogger}
	err := bot.ConfigureFromJson(groupmeConfigFileName)

	checkErr(err)

	gb.Bot = &bot
	gb.LoadHooks(botConfigFileName)
	//bot.AddHook("regex", funcion)

	log.Printf("Listening on %v...\n", bot.Server)
	http.HandleFunc("/", bot.Handler())
	log.Fatal(http.ListenAndServe(bot.Server, nil))
}

func genericResponse(msg groupmebot.InboundMessage) string {
	return "Got it!"
}
