package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/yut-kt/goholiday"
)

func main() {
	fromFlag := flag.String("from", "", "input from day yyyy/mm/dd")
	flag.Parse()
	from := *fromFlag

	t, _ := time.Parse("2006/01/02", from)
	now := time.Now()

	export := fmt.Sprintf("./%s.csv", "result")
	wFile, _ := newFile(export)
	writer := bufio.NewWriter(wFile)
	for {
		isHoliday := "0"
		thisWeek := t.Weekday()

		if goholiday.IsHoliday(t) {
			fmt.Printf("%s", t)
			isHoliday = "1"
		}
		if thisWeek == 6 || thisWeek == 0 {
			isHoliday = "1"
		}

		output := fmt.Sprintf("%s, %s, %s\n", t.Format("2006/1/2"), t.Weekday(), isHoliday)
		writer.WriteString(output)
		writer.Flush()

		t = t.Add(time.Hour * 24)
		if now.Sub(t) < 0 {
			break
		}
	}
	//fmt.Println(time.Now())
}

func newFile(fn string) (*os.File, bool) {
	_, exist := os.Stat(fn)
	fp, err := os.OpenFile(fn, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return fp, os.IsNotExist(exist)
}
