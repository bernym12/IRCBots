package dad

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/adammohammed/groupmebot"
)

// Run runs
func (gb *GroupMeBot) Run(botConfigFileName string) {
	rand.Seed(time.Now().Unix())

	lg := groupmebot.CSVLogger{LogFile: "bot.csv"}
	stdout := groupmebot.StdOutLogger{}
	combinedLogger := groupmebot.CompositeLogger{Loggers: []groupmebot.Logger{lg, stdout}}

	bot := groupmebot.GroupMeBot{Logger: combinedLogger}
	err := bot.ConfigureFromJson("dad_groupme_cfg.json")

	checkErr(err)

	gb.LoadHooks(botConfigFileName)
	//bot.AddHook("regex", funcion)

	log.Printf("Listening on %v...\n", bot.Server)
	http.HandleFunc("/", bot.Handler())
	log.Fatal(http.ListenAndServe(bot.Server, nil))
}
