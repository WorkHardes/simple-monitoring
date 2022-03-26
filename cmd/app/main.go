package main

import (
	"fmt"
	"log"

	g "github.com/gosnmp/gosnmp"
)

type WorkTime struct {
	Days    uint32
	Hours   uint32
	Minutes uint32
	Seconds uint32
}

func NewWorkTime() *WorkTime {
	return &WorkTime{}
}

func GetWorkTime(oids []string) (*WorkTime, error) {
	result, err := g.Default.Get(oids)
	if err != nil {
		return nil, fmt.Errorf("g.Default.Get(oids) err: %v;", err)
	}
	timeTicks := result.Variables[0].Value
	workTimeTicks := timeTicks.(uint32)
	workSeconds := workTimeTicks / 100

	wt := NewWorkTime()
	wt.Days = workSeconds / 86400
	wt.Hours = workSeconds / 3600
	wt.Minutes = workSeconds % 3600 / 60
	wt.Seconds = workSeconds % 216000 % 60

	return wt, nil
}

func main() {
	g.Default.Target = "192.168.1.1"
	if err := g.Default.Connect(); err != nil {
		log.Fatalf("Connect() err: %v", err)
	}
	defer g.Default.Conn.Close()

	oids := []string{"1.3.6.1.2.1.1.3.0"}
	workTime, err := GetWorkTime(oids)
	if err != nil {
		log.Fatalf("GetWorkTime() err: %v", err)
		return
	}
	fmt.Println("WorkTime:", workTime)
}
