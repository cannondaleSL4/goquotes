package scheduler

import (
	"flag"
	"fmt"
	"github.com/Syfaro/telegram-bot-api"
	tinkoff "github.com/TinkoffCreditSystems/invest-openapi-go-sdk"
	"github.com/goquotes/analyse"
	"github.com/goquotes/constants"
	"github.com/goquotes/controller"
	"github.com/jasonlvhit/gocron"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var tokenTg = flag.String("tokenTg", os.Getenv("TOKEN_TELEGRAM"), "your token")
var chatId = flag.String("chatId", os.Getenv("CHATID"), "your chatId")

func Scheduler() {
	gocron.Every(10).Minutes().Do(ping)

	gocron.Every(2).Minutes().Do(telega)

	gocron.Every(1).Monday().At("09:00").Do(telega)
	gocron.Every(1).Tuesday().At("09:00").Do(telega)
	gocron.Every(1).Wednesday().At("09:00").Do(telega)
	gocron.Every(1).Thursday().At("09:00").Do(telega)
	gocron.Every(1).Friday().At("09:00").Do(telega)
	gocron.Every(1).Friday().At("09:00").Do(telega)

	<-gocron.Start()
}

//this method for heroku only
func ping() {
	fmt.Println("ping ping ping")
	_, _ = http.Get("https://gogo-quotes.herokuapp.com/")
}

func telega() {
	resultDJ := task(constants.DOWJONES)
	log.Println("Analyse job for DJ telegram has been executed")
	sentResult(*resultDJ)
	resultRUS := task(constants.RUS)
	sentResult(*resultRUS)
	log.Println("Analyse job for Rus telegram has been executed")
}

func task(instr string) *[]analyse.AnalyzeResponse {
	var result *[]analyse.AnalyzeResponse
	var arrayOfCandle *[][]tinkoff.Candle
	fromTime := time.Now()
	arrayOfCandle, _ = controller.GetCandle(fromTime.AddDate(0, 0, -364), instr, tinkoff.CandleInterval1Day)
	result = analyse.GetAnalyse(arrayOfCandle, tinkoff.CandleInterval1Day)
	return result
}

func sentResult(result []analyse.AnalyzeResponse) {
	fmt.Println("*******************************************")
	fmt.Println("start telega")
	fmt.Printf("tokenTg: %s\n", *tokenTg)
	fmt.Printf("chatId: %s\n", *chatId)
	fmt.Println("*******************************************")

	bot, err := tgbotapi.NewBotAPI(*tokenTg)
	if err != nil {
		log.Printf("Some problem with telegram API %s\n", err)
		return
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	//var ucfg tgbotapi.UpdateConfig = tgbotapi.NewUpdate(0)
	//ucfg.Timeout = 60
	//updates, err := bot.GetUpdatesChan(ucfg)
	//
	//_ = updates

	//for update := range updates {
	//	if update.Message == nil {
	//		continue
	//	}
	//}

	chatIdInt64, err := strconv.ParseInt(*chatId, 10, 64)
	if err != nil {
		log.Printf("Err! Could not get chat id. err: %s", err)
		return
	}
	if len(result) == 0 {
		msg := tgbotapi.NewMessage(chatIdInt64, fmt.Sprintln("There are no any results for trading"))
		bot.Send(msg)
	} else {
		//msg := tgbotapi.NewMessage(chatIdInt64, fmt.Sprintln("There are results of analyse:"))
		//bot.Send(msg)
		strResult := "There are results of analyse: \n" + parseResult(&result)
		//msg1 := tgbotapi.NewMessage(chatIdInt64, fmt.Sprintln(result))
		msg1 := tgbotapi.NewMessage(chatIdInt64, strResult)
		bot.Send(msg1)
	}

	//for {
	//	select {
	//	case update := <-bot.Updates:
	//		// Пользователь, который написал боту
	//		UserName := update.Message.From.UserName
	//
	//		// ID чата/диалога.
	//		// Может быть идентификатором как чата с пользователем
	//		// (тогда он равен UserID) так и публичного чата/канала
	//		ChatID := update.Message.Chat.ID
	//
	//		// Текст сообщения
	//		Text := update.Message.Text
	//
	//		log.Printf("[%s] %d %s", UserName, ChatID, Text)
	//
	//		// Ответим пользователю его же сообщением
	//		reply := Text
	//		// Созадаем сообщение
	//		msg := tgbotapi.NewMessage(ChatID, reply)
	//		// и отправляем его
	//		bot.SendMessage(msg)
	//	}
	//
	//}
}

func parseResult(result *[]analyse.AnalyzeResponse) string {
	var line string
	for _, element := range *result {
		line += fmt.Sprintf("Interval: %s Name: %s Result: %s \n", element.Interval, element.Name, element.Result)
	}
	return line
}

//func readChat(bot *BotAPI) {
//	for {
//		select {
//		case update := <-bot.Updates:
//			// Пользователь, который написал боту
//			UserName := update.Message.From.UserName
//
//			// ID чата/диалога.
//			// Может быть идентификатором как чата с пользователем
//			// (тогда он равен UserID) так и публичного чата/канала
//			ChatID := update.Message.Chat.ID
//
//			// Текст сообщения
//			Text := update.Message.Text
//
//			log.Printf("[%s] %d %s", UserName, ChatID, Text)
//
//			// Ответим пользователю его же сообщением
//			reply := Text
//			// Созадаем сообщение
//			msg := tgbotapi.NewMessage(ChatID, reply)
//			// и отправляем его
//			bot.SendMessage(msg)
//		}
//
//	}
//}
