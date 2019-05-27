package dbops

import (
	_ "github.com/go-sql-driver/mysql"
	"GolangWebCourseware/config"
	"github.com/go-xorm/xorm"
)

var (
	Engine *xorm.Engine
	Err error
)

func init()  {
	Engine, Err = xorm.NewEngine(config.BasicsConfig.DriverName, config.BasicsConfig.Dsn)
	if Err != nil {
		panic(Err.Error())
	}
	Engine.ShowSQL(true)
	SynchronousData()
}

func SynchronousData()  {
	err := Engine.Sync2(new(Account))
	if err != nil {
		panic(err.Error())
	}
}
