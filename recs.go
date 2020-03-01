package main

import (
	"fmt"
)

type Rec struct {
	Title          string
	Date           string
	UsedTimeInSecs int
}

var Recs = []Rec{}

type RecsIO interface {
	Open() error
	GetAll() []Rec
	Add(Rec)
	Close()
}

type RecsCSV struct {
	FileName   string
	FileObject *File
}

func (rc RecsCSV) Open() error {
	f, err := os.OpenFile(rc.FileName, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	} else {
		rc.FileObject = f
		return nil
	}
}

func (rc RecsCSV) GetAll() []Rec {
	return nil
}

func (rc RecsCSV) Add(r Rec) error {
	content := fmt.Sprintf("%s\t%s\t%d", r.Title, r.Date, r.UsedTimeInSecs)
	_, err = rc.FileObject.Write([]byte(content))
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (rc RecsCSV) Close() {
	rc.FileObject.Close()
}
