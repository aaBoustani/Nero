package main

import (
	"net/http"
	"fmt"
	"net/url"
	"strings"
	"bytes"
	"encoding/json"
	"log"
	"strconv"
)

type Command struct {
	Token string
	TeamID string
	TeamDomain string
	ChannelID string
	ChannelName string
	UserID string
	UserName string
	Command string
	Text string
	ResponseURL *url.URL
}

type Attachment struct {
	Pretext string `json:"pretext"`
	Text string `json:"text"`
}

// TODO Verify token each time
// TODO secure the keys
func Give(w http.ResponseWriter, r *http.Request) {
	u, err := parseRequest(r)

	if err != nil {
			http.Error(w, err.Error(), 400)
			return
	}

	s := strings.Split(u.Text, " ")
	user := s[0]
	i := 1
	amount := 1
	if len(s) > 1 {
		amount, err = strconv.Atoi(s[1])
		if err != nil {
			amount = 1
			i = 1
		} else {
			i = 2
		}
	}

	if amount > NEROLIMIT {
		w.Write([]byte("You're sending more than what is allowed."))
		return
	}

	reason := strings.Join(s[i:], " ")
	dbUser := strings.Replace(user, "@", "", 1)
	go AddNero(dbUser, amount)

	msg := fmt.Sprintf("*@%s* gave you *%d Nero*", u.UserName, amount)
	go sendMsg(msg, user, reason)

	sMsg := fmt.Sprintf("You gave *%d Nero* to *%s*", amount, user)
	go sendMsg(sMsg, u.ChannelID, reason)

	w.Write([]byte(""))
}

func GetScore(w http.ResponseWriter, r *http.Request) {
	u, err := parseRequest(r)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	amount, err := GetNero(u.UserName)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if amount < 0 {
		http.Error(w, "User not found", 400)
		return
	}

	w.Write([]byte(fmt.Sprintf("You have *%d Nero*", amount)))
}

func GetAllScores(w http.ResponseWriter, r *http.Request) {
	res := PrintAll()
	for e := range res {
		w.Write([]byte(fmt.Sprintf("%s %d\n", res[e].User, res[e].Amount)))
	}
}

func commandFromValues(v url.Values) (Command, error) {
	u, err := url.Parse(v.Get("response_url"))
	if err != nil {
		return Command{}, err
	}

	return Command{
		Token:       v.Get("token"),
		TeamID:      v.Get("team_id"),
		TeamDomain:  v.Get("team_domain"),
		ChannelID:   v.Get("channel_id"),
		ChannelName: v.Get("channel_name"),
		UserID:      v.Get("user_id"),
		UserName:    v.Get("user_name"),
		Command:     v.Get("command"),
		Text:        v.Get("text"),
		ResponseURL: u,
	}, nil
}

func parseRequest(r *http.Request) (Command, error) {
	err := r.ParseForm()
	if err != nil {
		return Command{}, err
	}
	return commandFromValues(r.Form)
}

func sendMsg(msg string, rec string, att string) {
	token := "xoxp-2311130354-288237133235-357870053315-338b16386c80446ce6ab4f0b62e83e8b"
	form := url.Values{}
	form.Add("text", msg)
	form.Add("channel", rec)
	form.Add("token", token)

	a := append([]Attachment{}, Attachment{ Text: att })
	out, err := json.Marshal(a)
	if err != nil {
		log.Fatal(err)
	}
	form.Add("attachments", string(out))

	http.Post("https://slack.com/api/chat.postMessage", "application/x-www-form-urlencoded", bytes.NewBuffer([]byte(form.Encode())))
}
