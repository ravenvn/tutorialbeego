package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

type User struct {
	Id   int
	Name string `orm:"size(100)"`
}

func init() {
	orm.RegisterDataBase("default", "postgres", "postgres://postgres:mickey02@/tutorial?sslmode=disable", 30)
	orm.RegisterModel(new(User))
}
