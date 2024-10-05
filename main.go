package main

import (
	"fmt"

	"github.com/alfredosa/GoDiscordBot/bot"
	"github.com/alfredosa/GoDiscordBot/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err)
		return
	}

	bot.Start()

	<-make(chan struct{})

	return
}
