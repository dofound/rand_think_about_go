package main

import (
	"database/sql"
	"fmt"
	"github.com/BurntSushi/toml"
	_ "github.com/go-sql-driver/mysql"  //must import
)

func main() {

	//Get the my.conf information
	type dbConf struct {
		Addr       string `toml:"addr"`
		DriverName string `toml:"driver_name"`
	}
	type myConf struct {
		Dbconf dbConf `toml:"database"`
	}
	var myconf myConf
	_, err := toml.DecodeFile("my.conf", &myconf)
	if err != nil {
		fmt.Printf("decode err:%v", err)
	}
	fmt.Printf("conf :%v", myconf)

	//connect db
	db, err := sql.Open(myconf.Dbconf.DriverName, myconf.Dbconf.Addr)
	if err != nil {
		fmt.Printf("error:%v", err)
	}
	defer db.Close()
	fmt.Printf("db:%v", db)

	//get db info
	sqlQeury := "select * from my limit 1"
	row := db.QueryRow(sqlQeury)
	var (
		id, age int64
		name    string
	)
	err = row.Scan(&id, &age, &name)
	if err != nil {
		fmt.Printf("scan:%v", err)
	}
	fmt.Printf("data:%d,%d,%s", id, age, name)
}
