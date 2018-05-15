package main

import (
	"database/sql"
	"fmt"
	"github.com/BurntSushi/toml"
	_ "github.com/go-sql-driver/mysql" //must import
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
	//fmt.Printf("conf :%v", myconf)

	//connect db
	db, err := sql.Open(myconf.Dbconf.DriverName, myconf.Dbconf.Addr)
	if err != nil {
		fmt.Printf("error:%v", err)
	}
	defer db.Close()
	//fmt.Printf("db:%v", db)

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

	//获取数据
	sqlPuy := "select * from my"
	rowp,_:=db.Query(sqlPuy)
	printResult(rowp)

	//获取字段
	fieldQeury := "desc my"
	rowd,_ := db.Query(fieldQeury)
	printResult(rowd)
}

//数据扫描，返回map格式
func printResult(query *sql.Rows) {
	column, _ := query.Columns()              //读出查询出的列字段名
	values := make([][]byte, len(column))     //values是每个列的值，这里获取到byte里
	scans := make([]interface{}, len(column)) //因为每次查询出来的列是不定长的，用len(column)定住当次查询的长度
	for i := range values {                   //让每一行数据都填充到[][]byte里面
		scans[i] = &values[i]
	}
	results := make(map[int]map[string]string) //最后得到的map
	i := 0
	for query.Next() { //循环，让游标往下移动
		if err := query.Scan(scans...); err != nil { //query.Scan查询出来的不定长值放到scans[i] = &values[i],也就是每行都放在values里
			fmt.Println(err)
			return
		}
		row := make(map[string]string) //每行数据
		for k, v := range values {     //每行数据是放在values里面，现在把它挪到row里
			key := column[k]
			row[key] = string(v)
		}
		results[i] = row //装入结果集中
		i++
	}

	for k, v := range results { //查询出来的数组
		fmt.Printf("%d,%v\n",k, v)
	}

}
