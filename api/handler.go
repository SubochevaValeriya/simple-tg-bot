package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"tgBot/clients/telegram"
)

var bot *telegram.Bot

func init() {
	_, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	bot = telegram.NewBot(http.Client{}, os.Getenv("BOT_TOKEN"))
}

func Handler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Println(err)
	}
	var update telegram.UpdateResult

	err = json.Unmarshal(body, &update)
	fmt.Println(update)
	fmt.Println(body)
	bot.SendMessage(context.Background(), update)
}
