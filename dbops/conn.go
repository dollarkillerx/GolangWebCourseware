package dbops

import (
	"database/sql"
	"encoding/json"
	"os"
)

type dbConfig struct {
	DriverName string `json:"driverName"`
	Dsn string `json:"dsn"`
}

var (
	ConnDb *sql.DB
	err error
)

func init()  {
	config := getDbConfig()
	//建立conn
	driverName := config.DriverName //选择数据库引擎 这个我们用的是mysql
	dsn := config.Dsn //这个是dns  用户名:密码@(ip:端口)/数据库?charset=utf8
	ConnDb, err = sql.Open(driverName, dsn)
	if err != nil{
		panic(err.Error())
	}
}

// 获取database配置
func getDbConfig() *dbConfig {
	filePath := "./dbconfig.json"
	file, e := os.Open(filePath)
	defer file.Close()
	if e != nil {
		panic(e.Error())
	}

	config := &dbConfig{}
	decoder := json.NewDecoder(file)
	e = decoder.Decode(config)
	if e != nil {
		panic(e)
	}
	return config
}