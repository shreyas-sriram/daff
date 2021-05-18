package daff

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"gopkg.in/yaml.v2"
)

//	challenges:
//		- name:
//			url:
//			request:
//				method:
//				headers:
//					-
//					-
//				cookies:
//					-
//					-
//				body:
//			response:
//				status:
//		- name:
//			url:
//			request:
//				method:
//				headers:
//					-
//					-
//				cookies:
//					-
//					-
//				body:
//			response:
//				status:

type Config struct {
	Challenges map[string]Challenge `yaml:"challenges,omitempty"`
}

type Challenge struct {
	URL      string   `yaml:"url,omitempty"`
	Request  Request  `yaml:"request,omitempty"`
	Response Response `yaml:"response,omitempty"`
}

type Request struct {
	Method  string   `yaml:"method,omitempty"`
	Headers []string `yaml:"headers,omitempty"`
	Cookies []string `yaml:"cookies,omitempty"`
	Body    string   `yaml:"body,omitempty"`
}

type Response struct {
	Status int `yaml:"status,omitempty"`
}

const (
	headerDelimiter = ":"

	prefix = "!daff" // TBD, give a nice name

	responseMessage = "Challenge `%v` is %s\n"

	up   = ":thumbsup:"
	down = ":thumbsdown:"

	connectionRefused = "connection refused"
)

func New(file string) (*Config, error) {
	config := &Config{}

	err := config.parseFile(file)
	if err != nil {
		log.Printf("Error parsing config file: %v", err)
		return nil, err
	}

	return config, nil
}

// Print pretty print the config
func (c *Config) Print() {
	bytes, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Printf("%s\n", string(bytes))
}

// This function handles new message from channels that the bot has access to
func (c *Config) MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Check if the message is intended for the bot
	if !strings.HasPrefix(m.Content, prefix) {
		return
	}

	parts := strings.Split(m.Content, " ")
	if len(parts[1]) == 0 {
		return
	}

	challenge := parts[1]

	res, err := c.CheckSanity(challenge)
	if err != nil {
		log.Printf("Failed to check health: %v", err)
		if !strings.Contains(err.Error(), connectionRefused) {
			return
		}
	}

	var message string

	if res {
		message = fmt.Sprintf(responseMessage, challenge, up)
	} else {
		message = fmt.Sprintf(responseMessage, challenge, down)
	}

	log.Printf("Responding with message: %v", message)
	s.ChannelMessageSend(m.ChannelID, message)
}

// CheckSanity checks the health of a challenge
func (c *Config) CheckSanity(name string) (bool, error) {
	challenge, ok := c.Challenges[name]
	if !ok {
		log.Print("Challenge configuration not found")
		return false, fmt.Errorf("challenge configuration not found")
	}

	log.Printf("Challenge configuration found: %+v\n", challenge)

	req, err := createRequest(challenge)
	if err != nil {
		log.Printf("Error creating request: %v\n", err)
		return false, err
	}

	log.Printf("Request created: %+v\n", req)

	res, err := sendRequest(req)
	if err != nil {
		log.Printf("Error sending request: %v\n", err)
		return false, err
	}

	log.Printf("Received response: %+v\n", res)

	return isResponseValid(res, &challenge.Response), nil
}

// parseFile parses the config file into a struct
func (c *Config) parseFile(file string) error {
	fileBytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Printf("Error reading config file: %v", err)
		return err
	}

	log.Printf("Configuration file found\n")

	err = yaml.Unmarshal(fileBytes, c)
	if err != nil {
		log.Printf("Error unmarshalling config file: %v", err)
		return err
	}

	return nil
}
