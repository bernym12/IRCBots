package main

import dad "github.com/AFTERWAKE/IRCBots/dad/dadbot"

func main() {
	// dad.Irc.Run("irc_config.json", "dad_config.json", "muted_list.json")
	dad.GroupMe.Run("dad_groupme_config.json", "dad_config.json")
}
