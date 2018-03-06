package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	discordbot "github.com/ltsnuggie/discord-bot"
)

type Config struct {
	Mysql        SqlConfig `json:"mysql"`
	TestChannel  string    `json:"test_channel"`
	LogChannel   string    `json:"log_channel"`
	Token        string    `json:"token"`
	PollInterval int       `json:"poll_interval"`
}

// var db MysqlConnector

const configFilename = "./bot.conf.js"

func main() {

	conf, err := loadConfig()
	if err != nil {
		fmt.Println("Configfile ", configFilename, " not found...")
	}
	bot := discordbot.New(conf.Token)
	//load config
	bot.SetTestChannel(conf.TestChannel)
	bot.SetLogChannel(conf.LogChannel)
	bot.SendLogMessage("Bot starting up...")

	defer bot.Close()

	db := NewDBWrapper(conf)
	defer db.Close()

	// report := NewReport(db)

	poller := NewPoller(conf.PollInterval, &db)
	db.ShowTables()
	poller.Start()
	defer poller.Stop()

	commander := NewCommander(&db)

	bot.IsCaseSensative(false)
	bot.AddCommand("who am i", commander.WhoAmI)
	bot.AddCommand("i am", commander.IAm)
	bot.AddCommand("exit", commander.Exit)
	bot.AddCommand("echo", commander.Echo)
	// bot.AddCommand("my stats", MyStats)
	bot.AddCommand("what is my", commander.ReportStat)
	for {
	}

}

func loadConfig() (Config, error) {
	conf := Config{}
	file, err := ioutil.ReadFile(configFilename)
	if err != nil {
		//error reading file
		return conf, errors.New("Error reading file " + configFilename)
	}
	err = json.Unmarshal(file, &conf)
	if err != nil {
		//error unmarshalling
		return conf, errors.New("Error unmarshalling " + configFilename)
	}
	return conf, nil
}
