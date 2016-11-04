package main

import "github.com/parnurzeal/gorequest"
import "github.com/bitly/go-simplejson"
import "encoding/json"
import "fmt"

type Ticket struct {
	Acstatus     string `json:"acstatus"`
	Addtime      string `json:"addtime"`
	AuctionUrl   string `json:"auction_url"`
	CommonUrl    string `json:"common_url"`
	Cstatus      string `json:"cstatus"`
	Id           string `json:"id"`
	Oversea      string `json:"oversea"`
	Pretime      int64  `json:"pretime"`
	Special      string `json:"special"`
	Sstatus      string `json:"sstatus"`
	StandUrl     string `json:"stand_url"`
	Style        int64  `json:"style"`
	TeamId       string `json:"team_id"`
	Teamname     string `json:"teamname"`
	Theme        string `json:"theme"`
	TicketsPrice string `json:"tickets_price"`
	Title        string `json:"title"`
	Type         string `json:"type"`
	VipUrl       string `json:"vip_url"`
	Vstatus      string `json:"vstatus"`
}

func queryTickets(team, date string) {
	if date == "" {
		date = now("2006-01")
	}

	team = mapTeam(team)

	_, body, err := gorequest.
		New().
		Set("User-Agent", userAgent).
		Post("http://www.snh48.com/ticketsinfo.php").
		Query("act=choose").
		Send("team=" + team + "&date=" + date).
		End()

	if err != nil {
		fmt.Errorf("query ticket info error: %v", err)
	}

	data, _ := simplejson.NewJson([]byte(body))
	rows, _ := data.Array()

	for _, row := range rows {
		bytes, _ := json.Marshal(row)
		t := new(Ticket)

		err := json.Unmarshal(bytes, t)
		if err != nil {
			fmt.Errorf("map to ticket struct error: %v", err)
			continue
		}

		fmt.Println(t.Title)

		if t.Theme != "" {
			fmt.Println(t.Theme)
		}

		if t.Special != "" {
			fmt.Println(t.Special)
		}

		fmt.Printf("vip: %v\n", t.VipUrl)
		fmt.Printf("站: %v\n", t.StandUrl)
		fmt.Printf("普: %v\n", t.CommonUrl)

		fmt.Println("")
	}
}
