package main

import (
	_ "tutorial/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

// func main() {
// 	o := orm.NewOrm()

// 	user := User{Name: "slene"}

// 	// insert
// 	id, err := o.Insert(&user)
// 	fmt.Printf("ID: %d, ERR: %v\n", id, err)

// 	// update
// 	user.Name = "Chuột đồng"
// 	num, err := o.Update(&user)
// 	fmt.Printf("NUM: %d, ERR: %v\n", num, err)

// 	// read one
// 	u := User{Id: user.Id}
// 	err = o.Read(&u)
// 	fmt.Printf("ERR: %v\n", err)

// 	// delete
// 	// num, err = o.Delete(&u)
// 	// fmt.Printf("NUM: %d, ERR: %v\n", num, err)
// }
