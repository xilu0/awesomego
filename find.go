package main

import (
	"fmt"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"time"
)

type User struct {
	Id      int64
	Name    string
	Created time.Time `xorm:"created"`
	Update  time.Time `xorm:"updated"`
}

func main() {
	f := "find.db"
	os.Remove(f)

	orm, err := xorm.NewEngine("sqlite3", f)
	if err != nil {
		fmt.Println(err)
		return
	}
	orm.ShowSQL(true)

	err = orm.CreateTables(&User{})
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = orm.Insert(&User{Id: 1, Name: "xlw"})
	if err != nil {
		fmt.Println(err)
		return
	}

	users := make([]User, 0)
	err = orm.Find(&users)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(users)
}
