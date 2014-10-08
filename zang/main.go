package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"zang/models"
)

func init() {
	models.RegisteDb()
}

func main() {

	orm.RunSyncdb("default", false, true)
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库

	profile := new(models.Profile)
	profile.Age = 30

	user := new(models.User)
	user.Profile = profile
	user.Name = "slene"

	fmt.Println(o.Insert(profile))
	fmt.Println(o.Insert(user))

	/*beego.Router("/", &homecontroller{})
	beego.Run()*/
}
