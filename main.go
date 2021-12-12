package main

import (
	"fmt"
	"log"
	"main/lib"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

const (
	configFile = "config.toml"
)

func main() {
	c := lib.ReadConfig(configFile)

	dg, err := discordgo.New("Bot " + c.Token)
	if err != nil {
		log.Fatal(err)
	}

	// per Receive Message
	dg.AddHandler(checkMessage)

	err = dg.Open()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("BOT Running...")

	//Signal Wait
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	defer dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot == true {
		return
	}

	fmt.Println("< " + m.Content + " by " + m.Author.Username)
	if strings.Contains(m.Content, "www") {
		lib.MessageSendWithMention(s, m, "wwwww")
	}
}

func checkMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot == true {
		return
	}

	command := lib.ParseCommand(s, m)
	if command == true {
		return
	}

}
