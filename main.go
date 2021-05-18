package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/shreyas-sriram/ctf-health-bot/pkg/daff"
)

var (
	Token string
	file  string
)

func main() {
	flag.StringVar(&file, "f", "config.yaml", "Config file")
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()

	// Check token input
	if len(Token) == 0 {
		log.Printf("Please enter valid bot token\n\n")
		return
	}

	daff, err := daff.New(file)
	if err != nil {
		log.Printf("Error parsing config file: %v\n", err)
		return
	}

	log.Printf("Parsed configuration\n")
	daff.Print()

	// Create a new Discord session using the provided bot token
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		log.Printf("Error creating Discord session: %v\n", err)
		return
	}

	// Register the messageCreate function as a callback for MessageCreate events
	dg.AddHandler(daff.MessageCreate)

	// Set to receive message events
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening
	err = dg.Open()
	if err != nil {
		log.Printf("Error opening connection: %v\n", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received
	log.Printf("Bot is running\n")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session
	dg.Close()
}
