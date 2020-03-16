package scheduler

import (
	"flag"
	"github.com/jasonlvhit/gocron"
	"os"
)

var tokenTg = flag.String("tokenTg", os.Getenv("TOKEN_TELEGRAM"), "your token")

func Scheduler() {
	gocron.Every(1).Monday().At("09:00").Do(task)
	gocron.Every(1).Tuesday().At("09:00").Do(task)
	gocron.Every(1).Wednesday().At("09:00").Do(task)
	gocron.Every(1).Thursday().At("09:00").Do(task)
	gocron.Every(1).Friday().At("09:00").From(gocron.NextTick()).Do(task)
	<-gocron.Start()
}

//func task() {
//	var resultDj *[]analyse.AnalyzeResponse
//	//var resultRus *[]analyse.AnalyzeResponse
//	var arrayOfCandleDj *[][]tinkoff.Candle
//	//var arrayOfCandleRus *[][]tinkoff.Candle
//	fromTime := time.Now()
//	arrayOfCandleDj, _ = controller.GetCandle(fromTime.AddDate(0, 0, -364), constants.DOWJONES, tinkoff.CandleInterval1Day)
//	//arrayOfCandleRus, _ = controller.GetCandle(fromTime.AddDate(0, 0, -364), constants.RUS, tinkoff.CandleInterval1Day)
//	resultDj = analyse.GetAnalyse(arrayOfCandleDj, tinkoff.CandleInterval1Day)
//	//resultRus = analyse.GetAnalyse(arrayOfCandleRus, tinkoff.CandleInterval1Day)
//	_ = resultDj
//	_ = arrayOfCandleDj
//	//_ = resultRus
//	//_ = arrayOfCandleRus
//}

func task() {
	//fmt.Sprintln("dsds")
	//bot, err := tgbotapi.NewBotAPI(*tokenTg)
	//if err != nil {
	//	log.Panic(err)
	//}
	//
	//bot.Debug = true
	//log.Printf("Authorized on account %s", bot.Self.UserName)
	//
	////bot.NewMessage("@dmba_forecast",fmt.Sprintf("Hello, %s!", "Epta"))
	//msg := tgbotapi.NewMessage(356316908,fmt.Sprintf("Hello, %s!", "Epta"))
	//bot.SendMessage(msg)

}
