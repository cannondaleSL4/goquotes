package scheduler

import (
	"fmt"
	"gopkg.in/robfig/cron.v3"
	"time"
)

func Scheduler() {
	//go func() {
	//	oneF()
	//}()
	//
	//go func() {
	//	twoF()
	//}()
	oneF()
	twoF()
}

func oneF() {
	//ticker := time.NewTicker(time.Second)
	//defer ticker.Stop()
	//done := make(chan bool)
	//go func() {
	//	time.Sleep(5 * time.Second)
	//	done <- true
	//}()
	//for {
	//	select {
	//	case <-done:
	//		fmt.Println("1!")
	//		return
	//	case t := <-ticker.C:
	//		fmt.Println("Current time: ", t)
	//	}
	//}
}

func twoF() {

	c := cron.New()
	c.AddFunc("CRON_TZ=Europe/Moscow 05 10 * * * *", Test)
	c.AddFunc("CRON_TZ=Europe/Moscow 05 14 * * * *", Test)
	c.AddFunc("CRON_TZ=Europe/Moscow 05 18 * * * *", Test)
	c.AddFunc("CRON_TZ=Europe/Moscow 05 22 * * * *", Test)

	c.AddFunc("CRON_TZ=Europe/Moscow 04 17 * * * *", Test)
	c.AddFunc("CRON_TZ=Europe/Moscow 05 17 * * * *", Test)
	c.AddFunc("CRON_TZ=Europe/Moscow 06 17 * * * *", Test)
	c.AddFunc("CRON_TZ=Europe/Moscow 07 17 * * * *", Test)

	c.AddFunc("CRON_TZ=Europe/Moscow 40 17 * * * *", func() {
		fmt.Println("17/40")
	})

	c.AddFunc("CRON_TZ=Europe/Moscow 40 16 * * * *", func() {
		fmt.Println("16/40")
	})

	c.AddFunc("CRON_TZ=Europe/Moscow 40 15 * * * *", func() {
		fmt.Println("15/40")
	})

	c.AddFunc("CRON_TZ=Europe/Moscow 40 14 * * * *", func() {
		fmt.Println("14/40")
	})

	c.AddFunc("CRON_TZ=Europe/Moscow 40 18 * * * *", func() {
		fmt.Println("18/40")
	})

	c.AddFunc("CRON_TZ=Europe/Moscow 40 19 * * * *", func() {
		fmt.Println("19/40")
	})

	//c.AddFunc("03 * * * *", Test)
	//c.AddFunc("04 * * * *", Test)
	//c.AddFunc("05 * * * *", Test)
	//
	//c.AddFunc("@every 0h02m", func() {
	//	fmt.Println("lalla")
	//})

	//fmt.Println("start")

	loc, _ := time.LoadLocation("Europe/Moscow")

	fmt.Println(time.Now().In(loc).String())
	c.Start()

	//fmt.Printf("LALALLALA")
	//
	//ticker := time.NewTicker(10 * time.Second)
	//done := make(chan bool)
	//
	//go func() {
	//	for {
	//		select {
	//		case <-done:
	//			return
	//		case t := <-ticker.C:
	//			fmt.Println("Tick at", t)
	//		}
	//	}
	//}()
}

func Test() {
	fmt.Printf(time.Now().String())
}
