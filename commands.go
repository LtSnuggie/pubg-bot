package main

import (
	"os"
	"strconv"
	"strings"

	discordbot "github.com/ltsnuggie/discord-bot"
)

type Commander struct {
	Wrapper *DBWrapper
}

func NewCommander(db *DBWrapper) Commander {
	c := Commander{}
	c.Wrapper = db
	return c
}

func (c *Commander) WhoAmI(b *discordbot.Bot, args string) {
	id := b.GetMessageAuthorID()
	tag, err := c.GetGamerTag(id)
	if err != nil {
		b.Error(err)
	}
	b.SendMessage(tag)
}

func (c *Commander) IAm(b *discordbot.Bot, args string) {
	id := b.GetMessageAuthorID()
	err := c.SetGamerTag(id, args)
	if err != nil {
		b.Error(err)
	}
}

func (c *Commander) Exit(b *discordbot.Bot, args string) {
	// b.Close()
	// db.Close()
	os.Exit(0)
}

func (c *Commander) Echo(b *discordbot.Bot, args string) {
	b.SendMessage(args)
}

// func MyStats(b *discordbot.Bot, args string) {
// 	id := b.GetMessageAuthorID()
// 	tag, err := GetGamerTag(id)
// 	if err != nil {
// 		b.Error(err)
// 	}
//
// 	b.SendMessage(ReportStats(FetchUserDetails(tag)))
// }

func (c *Commander) ReportStat(b *discordbot.Bot, args string) {
	id := b.GetMessageAuthorID()
	tag, err := c.GetGamerTag(id)
	if err != nil {
		b.Error(err)
	}
	ud := FetchUserDetails(tag)
	for _, stats := range ud.Stats {
		for _, stat := range stats.Data {
			if strings.ToLower(stat.Name) == strings.ToLower(args) {
				b.SendMessage(strconv.FormatInt(int64(stat.Value), 10))
				return
			}
		}
	}
	b.SendMessage("No stat found for '" + args + "'")
}

func (c *Commander) GetGamerTag(discordID string) (string, error) {
	return "", nil
	// return db.GetUserGamertag(discordID), nil
}

func (c *Commander) SetGamerTag(discordID, gamertag string) error {
	return nil
	// return db.SetUserGamertag(discordID, gamertag), nil
}
