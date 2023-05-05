package telegram

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strings"
	randoms "tgBot/pkg"
	"time"
)

type UpdateResult struct {
	UpdateID int `json:"update_id"`
	Message  struct {
		MessageID int `json:"message_id"`
		From      struct {
			ID           int    `json:"id"`
			IsBot        bool   `json:"is_bot"`
			FirstName    string `json:"first_name"`
			Username     string `json:"username"`
			LanguageCode string `json:"language_code"`
		} `json:"from"`
		Chat struct {
			ID        int    `json:"id"`
			FirstName string `json:"first_name"`
			Username  string `json:"username"`
			Type      string `json:"type"`
		} `json:"chat"`
		Date int    `json:"date"`
		Text string `json:"text"`
	} `json:"message"`
}

type update struct {
	Ok     bool           `json:"ok"`
	Result []UpdateResult `json:"result"`
}

type Bot struct {
	httpClient http.Client
	token      string
	period     time.Duration
	rand       *rand.Rand
}

func NewBot(httpClient http.Client, token string) *Bot {
	return &Bot{httpClient: httpClient, token: token, period: 500 * time.Millisecond, rand: rand.New(rand.NewSource(time.Now().UnixNano()))}
}

func (b Bot) Start(ctx context.Context) error {
	lastUpdateId := 0
	for {
		select {
		case <-time.Tick(b.period):
			req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates?offset=%d", b.token, lastUpdateId+1), nil)
			if err != nil {
				log.Println("bot getUpdates error:", err)
				continue
			}
			response, err := b.httpClient.Do(req)
			if err != nil {
				log.Println("bot getUpdates error:", err)
				continue
			}
			if response.StatusCode != 200 {
				bytes, _ := io.ReadAll(response.Body)
				_ = response.Body.Close()
				log.Println("bot getUpdates error:", string(bytes))
				continue
			}
			var update update
			err = json.NewDecoder(response.Body).Decode(&update)
			if err != nil {
				log.Println("bot getUpdates error:", err)
				_ = response.Body.Close()
				continue
			}
			_ = response.Body.Close()

			fmt.Printf("update='%+v'\n", update)
			for _, res := range update.Result {
				b.SendMessage(ctx, res)
			}
			if len(update.Result) > 0 {
				lastUpdateId = update.Result[0].UpdateID
			}
		case <-ctx.Done():
			log.Println("bot getUpdates error:", ctx.Err())
			return nil
		}
	}
}

type keyboardButton struct {
}

type replyMarkup struct {
	Keyboard [][]string `json:"keyboard,omitempty"`
}

type message struct {
	ChatId      int         `json:"chat_id"`
	Text        string      `json:"text,omitempty"`
	ReplyMarkup replyMarkup `json:"reply_markup,omitempty"`
	Photo       string      `json:"photo,omitempty"`
	Caption     string      `json:"caption,omitempty"`
}

func (b Bot) SendMessage(ctx context.Context, res UpdateResult) {

	var err error
	txt := strings.TrimSpace(res.Message.Text)
	var replyText, photoURL, caption string
	var keyboard [][]string
	keyboard = [][]string{
		{"Random fact"},
		{"Random activity"},
		{"Random answer"},
		{"Random dog"},
	}
	//if txt == "Random cat" {
	//	//photo, err := randoms.RandomCat()
	//	//if err != nil {
	//	//	fmt.Errorf("can't get random cat: %w", err)
	//	//}
	//	photoURL = "https://cataas.com/cat"
	//
	//}

	switch txt {
	case "/start":
		replyText = fmt.Sprintf("Привет, %s 👋", res.Message.From.FirstName)
	case "Random answer":
		answer, err := randoms.RandomAnswer()
		if err != nil {
			fmt.Errorf("can't get random answer: %w", err)
			break
		}
		photoURL = answer.Image
		caption = answer.Answer
	case "Random dog":
		photoURL, err = randoms.RandomDog()
		if err != nil {
			fmt.Errorf("can't get random dog: %w", err)
		}
	case "Random fact":
		replyText, err = randoms.RandomFact()
		if err != nil {
			fmt.Errorf("can't get random fact: %w", err)
		}
	case "Random activity":
		replyText = "Choose number of participants"
		keyboard = [][]string{
			{"1 participant"},
			{"More than 1 participant"},
		}

	case "1 participant":
		replyText, err = randoms.RandomActivity(1, 1)
		if err != nil {
			fmt.Errorf("can't get random activity: %w", err)
		}
	case "More than 1 participant":
		replyText, err = randoms.RandomActivity(2, 20)
		if err != nil {
			fmt.Errorf("can't get random activity: %w", err)
		}
	}

	msg := message{
		ChatId:      res.Message.Chat.ID,
		Text:        replyText,
		ReplyMarkup: replyMarkup{Keyboard: keyboard},
		Photo:       photoURL,
		Caption:     caption,
	}

	byt, err := json.Marshal(msg)
	if err != nil {
		log.Println("bot sendMessage error:", err)
		return
	}
	fmt.Printf("msg='%+v'\n", msg)
	var req *http.Request
	if msg.Photo != "" {
		req, err = http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("https://api.telegram.org/bot%s/sendPhoto", b.token), bytes.NewReader(byt))
		if err != nil {
			log.Println("bot sendMPhoto error:", err)
			return
		}
	} else {
		req, err = http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", b.token), bytes.NewReader(byt))
		if err != nil {
			log.Println("bot sendMessage error:", err)
			return
		}
	}

	req.Header.Add("Content-Type", "application/json")
	response, err := b.httpClient.Do(req)
	if err != nil {
		log.Println("bot sendMessage error:", err)
		return
	}
	if response.StatusCode != 200 {
		bytes, _ := io.ReadAll(response.Body)
		defer response.Body.Close()
		log.Println("bot sendMessage error:", string(bytes))
		return
	}
}
