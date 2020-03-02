package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Rec struct {
	Title          string
	Date           string
	UsedTimeInSecs int
}

func (r Rec) Format() string {
	mins := r.UsedTimeInSecs / 60
	secs := r.UsedTimeInSecs - 60*mins

	timeStr := ""
	if mins == 0 {
		timeStr = fmt.Sprintf("%d秒", secs)
	} else {
		timeStr = fmt.Sprintf("%d分%秒", mins, secs)
	}
	return fmt.Sprintf("%s\t%s\t%d\n", r.Title, r.Date, timeStr)
}

type RecsIO interface {
	Init() error
	Add(Rec)
	Query(title string) []Rec
}

type RecsCSV struct {
	FileName string
	Data     map[string]([]Rec)
}

func (rc RecsCSV) appendData(r Rec) {
	ra, ok := rc.Data[r.Title]
	if !ok {
		ra = []Rec{}
	}

	ra = append(ra, r)
	rc.Data[r.Title] = ra
}

func (rc RecsCSV) Init() error {
	//rc.Data = make(map[string]([]Rec))

	f, err := os.Open(rc.FileName)
	if err != nil {
		return err
	}
	defer f.Close()

	br := bufio.NewReader(f)
	for {
		s, _, e := br.ReadLine()
		if e == io.EOF {
			break
		}
		sa := strings.SplitN(string(s), "\t", 3)
		if len(sa) < 3 {
			continue
		}

		rec := Rec(sa[0], sa[1], sa[2])
		rc.appendData(rec)
	}
	return nil
}

func (rc RecsCSV) Query(title string) []Rec {
	ra, ok := rc.Data[title]
	if ok {
		return ra
	} else {
		return nil
	}
}

func (rc RecsCSV) Add(r Rec) error {
	rc.appendData(r)

	f, err := os.OpenFile(rc.FileName, os.O_WRONLY, 0666)
	if err != nil {
		return error
	}
	defer f.Close()

	_, err = f.Write([]byte(r.Format()))
	if err != nil {
		return error
	}
}
