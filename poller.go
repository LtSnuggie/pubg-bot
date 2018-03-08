package main

import (
	"fmt"
	"strconv"
	"time"
)

type Poller struct {
	WatchList    []string
	Ticker       *time.Ticker
	Done         chan bool
	PollInterval time.Duration
	Wrapper      DBWrapper
}

func NewPoller(interval int, db *DBWrapper) Poller {
	p := Poller{}
	p.PollInterval = time.Duration(interval)
	if interval < 1 {
		p.PollInterval = 30
	}
	p.Wrapper = *db
	p.Done = make(chan bool)
	return p
}

func (p *Poller) Start() {
	fmt.Println("Poller started...")
	p.Ticker = time.NewTicker(p.PollInterval * time.Second)
	p.WatchList = p.Wrapper.ListWatch()
	fmt.Println(p.WatchList)
	go func() {
		for {
			select {
			case <-p.Ticker.C:
				p.Poll()
			case <-p.Done:
				return
			}
		}
	}()
	time.Sleep(10 * time.Second)
}

func (p *Poller) Poll() {
	fmt.Println("Polling...")
	for _, tag := range p.WatchList {
		fmt.Println(tag)
		details := FetchUserDetails(tag)
		save := false
		if p.Wrapper.GetStats(tag) != "" {
			fetched, err := details.CheckStat("Matches Played")
			if err != nil {
				fmt.Println(err.Error())
			} else {
				db, err := strconv.ParseFloat(p.Wrapper.GetStat("Matches_Played", tag), 32)
				if err != nil {
					fmt.Println("No stats found for", tag)
				} else if float64(fetched) != db {
					save = true
				}
			}
		} else {
			save = true
		}
		if save {
			fmt.Println("New Stat detected...")
			if len(details.Stats) != 0 {
				p.Wrapper.SetStats(details)
			}
		}

	}
}

func (p *Poller) Stop() {
	p.Done <- true
}
