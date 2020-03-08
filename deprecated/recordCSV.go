package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type RecsCSV struct {
	FileName string
	Data     map[string]([]Record)
}

func (rc RecsCSV) appendData(r Record) {
	ra, ok := rc.Data[r.Key]
	if !ok {
		ra = []Record{}
	}

	ra = append(ra, r)
	rc.Data[r.Key] = ra
}

func (rc *RecsCSV) Load() error {
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

		rec := Record{
			Key:  sa[0],
			Date: 349484948,
			Secs: 0,
		}
		rc.appendData(rec)
	}
	return nil
}

func (rc *RecsCSV) Query(id string) []Record {
	ra, ok := rc.Data[id]
	if ok {
		return ra
	} else {
		return nil
	}
}

func (rc *RecsCSV) Add(r Record) error {
	rc.appendData(r)

	f, err := os.OpenFile(rc.FileName, os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write([]byte(fmt.Sprintf("%s\t%s\t%d\n", r.Key, r.Date, r.Secs)))
	if err != nil {
		return err
	}
	return nil
}
