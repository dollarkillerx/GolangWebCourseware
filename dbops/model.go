package dbops

type Account struct {
	Id int64 `xorm:"pk autoincr"`
	Name string `xorm:"unique"`
	Balance float64
	Version int `xorm:"version"` //乐观锁
}

