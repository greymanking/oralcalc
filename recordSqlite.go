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

func (rs *RecsSqlite) Query(user, key string) []Record {
	userId := rs.queryUser(user)
	if userId <= 0 {
		return nil
	}
	var recs []Record
	rows, err := rs.dbp.Query("SELECT key,date,secs FROM recs where user=? and key=?",
		userId, key)
	if err != nil {
		return nil
	}

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

func (rs *RecsSqlite) All(user string) []Record {
	userId := rs.queryUser(user)
	if userId <= 0 {
		return nil
	}
	var recs []Record
	rows, err := rs.dbp.Query("SELECT key,date,secs FROM recs where user=?",
		userId)
	if err != nil {
		return nil
	}

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

func (rs *RecsSqlite) Close() {
	rs.dbp.Close()
}

/*
func main() {
	var rsObject RecsSqlite
	rs := &rsObject
	if err := rs.Load(); err != nil {
		log.Panic(err)
	}
	defer rs.Close()

	var rec Record
	rec.Key = "001"
	rec.Date = 2000012230
	rec.Secs = 230

	//rs.Add("jinyinuo", rec)
	log.Print(rs.All("jinyinuo2"))
	log.Print(rs.Query("jinyinuo", "001"))
}

func mainsqlite() {
	db, err := sql.Open("sqlite3", "./foo.db")
	checkErr(err)

	//插入数据
	stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
	checkErr(err)

	res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
	//更新数据
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("astaxieupdate", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	//查询数据
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	//删除数据
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	db.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
*/
