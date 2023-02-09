// https://discordapp.com/oauth2/authorize?&client_id=1073291554342572092&scope=bot
package main

import (
	"fmt"

	"github.com/12A-r-p-i-t/discordbot/bot"
	"github.com/12A-r-p-i-t/discordbot/config"
)

func main() {
	err := config.ReadConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	bot.Start()
	<-make(chan struct{})
	return
}
