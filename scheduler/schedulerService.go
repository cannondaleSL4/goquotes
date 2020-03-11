package scheduler

import "github.com/jasonlvhit/gocron"

func Scheduler() {
	gocron.Every(1).Monday().At("10:00").Do(task)
	gocron.Every(1).Tuesday().At("10:00").Do(task)
	gocron.Every(1).Wednesday().At("10:00").Do(task)
	gocron.Every(1).Thursday().At("10:00").Do(task)
	gocron.Every(1).Friday().At("10:00").From(gocron.NextTick()).Do(task)
	<-gocron.Start()
}

func task() {
	//var result *[]analyse.AnalyzeResponse
	//var arrayOfCandle *[][]tinkoff.Candle
	//fromTime := time.Now()
	//arrayOfCandle, _ = GetCandle(fromTime.AddDate(0, 0, -364), constants.DOWJONES, tinkoff.CandleInterval1Day)
	//result = analyse.GetAnalyse(arrayOfCandle, tinkoff.CandleInterval1Day)
	//_ = result
	//_ = arrayOfCandle
}
