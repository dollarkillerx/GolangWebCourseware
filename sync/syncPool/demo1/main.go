/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-12
* Time: 上午10:50
* */
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"sync"
)

func main() {
	pool := &sync.Pool{
		New: func() interface{} {
			return mysql()
		},
	}

	db := pool.Get().(*sql.DB)
	ping := db.Ping()
	if ping != nil {
		panic(ping.Error())
	}else{
		fmt.Println("mysql OK")
	}
	pool.Put(db)
}

func mysql() *sql.DB {
	db, e := sql.Open("mysql", "one:2D7y4DEwYZfN43z3@(127.0.0.1:3306)/one?charset=utf8")
	if e != nil {
		panic(e.Error())
	}
	return db
}