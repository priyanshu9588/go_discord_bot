package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	// Create a new Discord session using the bot token.
	dg, err := discordgo.New("Bot " + BOT_TOKEN)
	if err != nil {
		fmt.Println("Error creating Discord session:", err)
		return
	}

	// Add a message handler for the "messageCreate" event.
	dg.AddHandler(messageCreate)

	// Open a websocket connection to Discord and begin listening for events.
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening connection:", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages from bots and our own messages.
	if m.Author.Bot || m.Author.ID == s.State.User.ID {
		return
	}

	// Respond to a user's message.
	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "pong")
	}
}
