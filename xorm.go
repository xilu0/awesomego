package main

import "os"
import "fmt"
import "github.com/go-xorm/xorm"
import (
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Id   int64
	Name string
}

type LoginInfo struct {
	Id        int64
	IP        string
	UserId    int64
	TimeStamp string `xorm: "<-"`
	Nonese    int    `xorm:"->"`
}

func main() {
	f := "singleMapping.db"
	os.Remove(f)

	orm, err := xorm.NewEngine("sqlite3", f)
	if err != nil {
		fmt.Println(err)
		return
	}
	orm.ShowSQL(true)
	orm.CreateTables(&User{}, &LoginInfo{})
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = orm.Insert(&User{1, "xiw"}, &LoginInfo{1, "127.0.0.1", 1, "", 23})
	if err != nil {
		fmt.Println(err)
		return
	}
	info := LoginInfo{}
	_, err = orm.ID(1).Get(&info)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(info)
}
