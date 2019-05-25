package dbops

import (
	"GolangWebCourseware/defs"
	"GolangWebCourseware/utils"
	_ "github.com/go-sql-driver/mysql" //注意啊这个要手动引入 (注:前面那个_就是只执行这个包的init方法)
	"log"
)

func RegisterUser(user *defs.User) error {
	name := user.Name
	password := user.Password

	uuid, _ := utils.NewUUIDSimplicity()

	password = utils.Md5String(password + uuid)

	//预编译插入sql 防止sql注入
	stmt, e := ConnDb.Prepare("INSERT INTO `user`(`user`,`password`,`salt`) VALUE (?,?,?)")
	defer stmt.Close() //延迟结束资源

	if e != nil {
		log.Println(e.Error())
		return e
	}

	_, e = stmt.Exec(name, password, uuid) //执行非查询的sql语句
	return e
}