package dad

import (
	"encoding/json"
	// "fmt"
	"io/ioutil"
	"os"
	// log "gopkg.in/inconshreveable/log15.v2"
)

func (gb *GroupMeBot) LoadHooks(botFileName string) {

}

// ReadConfig reads each config file and returns a struct of each
func (ib *IRCBot) ReadConfig(ircFileName, botFileName, mutedFileName string) {
	ib.IRCConfig = readIrcConfiguration(ircFileName)
	ib.BotConfig = readBotConfiguration(botFileName)
	ib.Muted = readMutedList(mutedFileName)
}

func readIrcConfiguration(fileName string) IRCConfiguration {
	file, _ := os.Open(fileName)
	defer file.Close()
	decoder := json.NewDecoder(file)
	config := IRCConfiguration{}
	err := decoder.Decode(&config)
	checkErr(err)
	return config
}

func readBotConfiguration(fileName string) BotConfiguration {
	file, _ := os.Open(fileName)
	defer file.Close()
	decoder := json.NewDecoder(file)
	config := BotConfiguration{}
	err := decoder.Decode(&config)
	checkErr(err)
	return config
}

func readMutedList(fileName string) MutedList {
	file, _ := os.Open(fileName)
	defer file.Close()
	decoder := json.NewDecoder(file)
	config := MutedList{}
	err := decoder.Decode(&config)
	checkErr(err)
	return config
}

// UpdateBotConfig parses the existing bot configuration data into
// a writable format and updates the corresponding data file.
func (ib *IRCBot) UpdateBotConfig() {
	jsonData, err := json.MarshalIndent(ib.BotConfig, "", "    ")
	checkErr(err)
	ioutil.WriteFile(ib.BotConfigFile, jsonData, 0644)
}
