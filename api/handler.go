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
	//err := bot.Start(ctx)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//var err error
	//bot, err = tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	//if err != nil {
	//	log.Panic(err)
	//}
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

	//var update tgbotapi.Update
	//
	//err = json.Unmarshal(body, &update)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//
	//ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	//defer stop()
	//
	//bot := mybot.NewBot(http.Client{}, os.Getenv("MYVKBOT_TOKEN"))
	//err := bot.Start(ctx)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//
	//log.Println("why")
	//
	//if update.Message != nil {
	//	// Construct a new message from the given chat ID and containing
	//	// the text that we received.
	//	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	//
	//	// If the message was open, add a copy of our numeric keyboard.
	//	switch update.Message.Text {
	//	case "open":
	//		msg.ReplyMarkup = firstLineButtons
	//
	//	}
	//
	//	// Send the message.
	//	if _, err = bot.Send(msg); err != nil {
	//		log.Fatal(err)
	//	}
	//}
	//
	//if update.Message.Text != "" {
	//	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
	//
	//	//telegram.SendMsg(update, bot)
	//}
	//
	//if update.CallbackQuery != nil {
	//	// Respond to the callback query, telling Telegram to show the user
	//	// a message with the data received.
	//	callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
	//	if _, err := bot.Request(callback); err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	// And finally, send a message containing the data received.
	//	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
	//	if _, err := bot.Send(msg); err != nil {
	//		log.Fatal(err)
	//	}
	//}
}

//var firstLineButtons = tgbotapi.NewInlineKeyboardMarkup(
//	tgbotapi.NewInlineKeyboardRow(
//		tgbotapi.NewInlineKeyboardButtonData("random fact", "random fact"),
//		tgbotapi.NewInlineKeyboardButtonData("random gif", "random gif"),
//		tgbotapi.NewInlineKeyboardButtonData("random cat", "random cat"),
//		tgbotapi.NewInlineKeyboardButtonData("random number", "random number"),
//	),
//)
//
//var secondLineButtons = tgbotapi.NewInlineKeyboardMarkup(
//	tgbotapi.NewInlineKeyboardRow(
//		tgbotapi.NewInlineKeyboardButtonData("lower number", "lower number"),
//		tgbotapi.NewInlineKeyboardButtonData("large number", "large number"),
//	),
//)
