package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Id   int64
	Name string
}

func main() {
	runtime.GOMAXPROCS(2)
	fmt.Println("-----start sqlite go routines")
	engine, err := sqliteEngine()
	if err != nil {
		fmt.Println(err)
		return
	}
	engine.ShowSQL(true)
	fmt.Println(engine)
	test(engine)
	fmt.Println("test end")
	engine.Close()
}

var u = &User{}

func test(engine *xorm.Engine) {
	err := engine.CreateTables(u)
	if err != nil {
		fmt.Println(err)
		return
	}

	size := 100
	queue := make(chan int, size)

	for i := 0; i < size; i++ {
		go func(x int) {
			err := engine.Ping()
			if err != nil {
				fmt.Println(err)
			} else {
				for j := 0; j < 10; j++ {
					if x+j < 2 {
						_, err = engine.Get(u)
					} else if x+j < 4 {
						users := make([]User, 0)
						err = engine.Find(&users)
					} else if x+j < 8 {
						_, err = engine.Count(u)
					} else if x+j < 16 {
						_, err = engine.Insert(&User{Name: "xlw"})
					} else if x+j < 32 {
						_, err = engine.ID(1).Delete(u)
					}
					if err != nil {
						fmt.Println(err)
						queue <- x
						return
					}
				}
				fmt.Printf("%v success!", x)
			}
			queue <- x
		}(i)
	}

	for i := 0; i < size; i++ {
		<-queue
	}

}

func sqliteEngine() (*xorm.Engine, error) {
	os.Remove("test.db")
	return xorm.NewEngine("sqlite3", "./goroutine.db")

}
