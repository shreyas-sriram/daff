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

// Variables used for command line parameters
var (
	Token string
	file  string
)

const (
	up   = "up"
	down = "down"
)

func main() {
	flag.StringVar(&Token, "t", "", "Bot Token") // ODQzODgwNjI0NDY4MzkzOTg0.YKKTEw.3XrUTJ4FcRk1ZJYldwOb8B2g860
	flag.StringVar(&file, "f", "config.yaml", "Config file")
	flag.Parse()

	// Check token input
	if len(Token) == 0 {
		log.Println("Please enter valid bot token")
		return
	}

	daff, err := daff.New(file)
	if err != nil {
		log.Printf("Error parsing config file: %v", err)
		return
	}

	log.Printf("Parsed configuration\n")
	daff.Print()

	// res, err := daff.CheckSanity("chall-1")
	// if err != nil {
	// 	log.Printf("Failed to check health: %v", err)
	// 	return
	// }

	// if res {
	// 	log.Printf("Challenge is: %+v\n", up)
	// } else {
	// 	log.Printf("Challenge is: %+v\n", down)
	// }

	// Create a new Discord session using the provided bot token
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		log.Printf("Error creating Discord session: %v", err)
		return
	}

	// Register the messageCreate function as a callback for MessageCreate events
	dg.AddHandler(daff.MessageCreate)

	// Set receiving message events
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening
	err = dg.Open()
	if err != nil {
		log.Printf("Error opening connection: %v", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received
	log.Println("Bot is running")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session
	dg.Close()
}
