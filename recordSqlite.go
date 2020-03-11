package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type RecsSqlite struct {
	dbp *sql.DB
}

//以下函数的接收者必须为RecsSqlite的指针，否则函数执行时只改变拷贝的dbp
//会导致后面一系列方法都不能正确访问到dbp这个指针
func (rs *RecsSqlite) Load() error {
	db, err := sql.Open("sqlite3", "db.sqlite")
	if err != nil {
		log.Panic(err)
		return err
	} else {
		rs.dbp = db
		return nil
	}
}

func (rs *RecsSqlite) queryUser(name string) int {
	var id int
	err := rs.dbp.QueryRow("SELECT id FROM users where name=?", name).Scan(&id)

	if err == sql.ErrNoRows {
		fmt.Printf("no user with name %s\n", name)
	} else if err != nil {
		log.Fatal(err)
		return -2
	}
	return id

}

func (rs *RecsSqlite) Add(user string, rec Record) error {
	userId := rs.queryUser(user)
	if userId <= 0 {
		return errors.New("user query failed")
	}
	stmt, err := rs.dbp.Prepare("INSERT INTO recs(user, key, date, secs) values(?,?,?,?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(userId, rec.Key, rec.Date, rec.Secs)
	if err != nil {
		return err
	}
	return nil
}

func RowsToRecords(rows *sql.Rows) []Record {
	var recs []Record

	for rows.Next() {
		var rec Record
		err := rows.Scan(&rec.Key, &rec.Date, &rec.Secs)
		if err != nil {
			log.Print(err)
		} else {
			recs = append(recs, rec)
		}
	}
	return recs
}

func (rs *RecsSqlite) Query(user, key string) []Record {
	userId := rs.queryUser(user)
	if userId <= 0 {
		return nil
	}

	rows, err := rs.dbp.Query("SELECT key,date,secs FROM recs where user=? and key=?",
		userId, key)
	if err != nil {
		return nil
	}

	return RowsToRecords(rows)
}

func (rs *RecsSqlite) All(user string) []Record {
	userId := rs.queryUser(user)
	if userId <= 0 {
		return nil
	}

	rows, err := rs.dbp.Query("SELECT key,date,secs FROM recs where user=?",
		userId)
	if err != nil {
		return nil
	}

	return RowsToRecords(rows)
}

func (rs *RecsSqlite) Close() {
	rs.dbp.Close()
}
