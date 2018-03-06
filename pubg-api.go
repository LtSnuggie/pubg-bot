package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

// type ResponsePayLoad struct {
// 	Data UserDetails `json:"Lt Snuggie"`
// }
type ResponsePayLoad map[string]UserDetails

type UserDetails struct {
	XUID        string   `json:"xuid"`
	Stats       []Stats  `json:"stats"`
	Friends     []Friend `json:"friends"`
	Clips       []Clip   `json:"clips"`
	Icon        string   `json:"icon"`
	LastUpdated string   `json:"lastupdated"`
	Gamertag    string
}

type Stats struct {
	Title string      `json:"title"`
	Data  []StatsData `json:"data"`
}

type StatsData struct {
	Name  string  `json:"name"`
	Value float32 `json:"value"`
}

type Friend struct {
	XUID     int    `json:"xuid"`
	GamerTag string `json:"gamertag"`
	Icon     string `json:"icon"`
}

type Clip struct {
	Thumbnail string `json:"thumbnail"`
	Link      string `json:"link"`
	Duration  int    `json:"duration"`
	Published string `json:"datePublished"`
	Recorded  string `json:"dateRecorded"`
}

func (d *UserDetails) CheckStat(stat string) (float32, error) {
	for _, group := range d.Stats {
		for _, data := range group.Data {
			if strings.EqualFold(data.Name, stat) {
				return data.Value, nil
			}
		}
	}
	return 0, errors.New("Stat '" + stat + "' not found")
}

func FetchUserDetails(user string) UserDetails {
	s := strings.Replace(user, " ", "%20", -1)
	url := "http://ec2-52-34-157-203.us-west-2.compute.amazonaws.com:3000/getProfileData/" + s
	// url := "http://pubgxboxstats.com/getProfileData/" + s
	fmt.Println(url)
	client := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal(err)
	}

	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	req.Header.Set("Accept-Encoding", "gzip, deflate")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		// log.Fatal(err)
		return UserDetails{}
	}
	var resp ResponsePayLoad
	json.NewDecoder(res.Body).Decode(&resp)
	ud := resp[user]
	ud.Gamertag = user
	return ud
}
