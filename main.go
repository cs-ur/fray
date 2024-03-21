package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	session, err := discordgo.New("Bot " + GetDotEnv("DISCORD_TOKEN"))
	GetError(err)

	session.AddHandler(MessageCreate)

	session.Open()
	defer session.Close()

	fmt.Println("Bot online!")
	// wait for stop
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func MessageCreate(discord *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == discord.State.User.ID {
		return
	}

	if message.Content == "ping" {
		discord.ChannelMessageSend(message.ChannelID, "pong!")
	}
}

func GetDotEnv(key string) string {
	err := godotenv.Load()
	GetError(err)

	return os.Getenv(key)
}

func GetError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
