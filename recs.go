package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Rec struct {
	ID             string
	Date           string
	UsedTimeInSecs int
}

type RecsIO interface {
	Init() error
	Add(Rec)
	Query(id string) []Rec
}

type RecsCSV struct {
	FileName string
	Data     map[string]([]Rec)
}

func (rc RecsCSV) appendData(r Rec) {
	ra, ok := rc.Data[r.ID]
	if !ok {
		ra = []Rec{}
	}

	ra = append(ra, r)
	rc.Data[r.ID] = ra
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

		rec := Rec{
			ID:             sa[0],
			Date:           sa[1],
			UsedTimeInSecs: 0,
		}
		rc.appendData(rec)
	}
	return nil
}

func (rc RecsCSV) Query(id string) []Rec {
	ra, ok := rc.Data[id]
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
		return err
	}
	defer f.Close()

	_, err = f.Write([]byte(fmt.Sprintf("%s\t%s\t%d\n", r.ID, r.Date, r.UsedTimeInSecs)))
	if err != nil {
		return err
	}
	return nil
}
