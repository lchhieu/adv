package main

import (
	_ "bee-project/routers"
	"github.com/astaxie/beego"
)

func init() {
}

func main() {
	/*fmt.Println("Start..")
	o := orm.NewOrm()
	o.Using("default") // Using default, you can use other database

	profile := new(Profile)
	profile.Age = 24
	profile.Address = "HCM"

	user := new(User)
	user.Profile = profile
	user.Name = "slene"

	// insert
	id, err := o.Insert(profile)
	if err == nil {
		fmt.Printf("ID: %v, ERR: %s, %v\n", id, err, profile)

		user.Profile = profile
		id, err = o.Insert(user)
		fmt.Printf("ID: %v, ERR: %s, %v\n", id, err, user)
	}*/

	/*var maps []orm.Params
	num, err := o.QueryTable("users").Values(&maps)
	if err == nil {
		fmt.Printf("Result Nums: %d\n", num)
		for _, m := range maps {
			fmt.Println(m["Id"], m["Name"])
		}
		}*/

	/*for _, m := range users {
	fmt.Println(m["Id"], m["Name"])
	}*/

	/*var index int64
	for index = 0; index < num; index++ {
		fmt.Println(users[index]["name"])
	}
	*/
	/*qs := o.QueryTable("users")

	for index := 0; index < len(qs); index++ {
		fmt.Println(item.name)
	}
	*/
	/**/
	beego.Run()
}
