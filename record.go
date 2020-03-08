package main

type Record struct {
	Key  string `json:"key" sqlite:"key"`
	Date uint   `json:"datems" sqlite:"date"`
	Secs int    `json:"secs" sqlite:"secs"`
}

type RecordIO interface {
	Load() error
	Add(user string, rec Record) error
	Query(user, key string) []Record
	All(user string) []Record
	Close()
}
