package main

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml" //配置文件解包
)

//MAIN
//注意 在读取配置文件时 各属性必须是公开的，也即是首字母须大写
func main() {
	type dbConf struct {
		Addr  string `toml:"addr"`
		Use   string `toml:"use"`
		Pass  string `toml:"pass"`
		Retry int    `toml:"retry"`
		Port  string `toml:"port"`
	}
	type memConf struct {
		Addr []string `toml:"addr"`
	}
	type serviceConf struct {
		Addr []string `toml:"addr"`
	}
	type redisConf struct {
		Addr []string `toml:"addr"`
	}
	type sysConfig struct {
		Dbconf      dbConf      `toml:"database"`
		Memconf     memConf     `toml:"memcache"`
		Serviceconf serviceConf `toml:"service"`
		Redisconf   redisConf   `toml:"redis"`
	}

	var sysconfig sysConfig
	var configPath string
	flag.StringVar(&configPath, "config", "my.conf", "server config.")
	flag.Parse()

	fmt.Println(configPath)
	if _, err := toml.DecodeFile(configPath, &sysconfig); err != nil {
		fmt.Printf("err:%v", err)
	}
	fmt.Printf("%v", sysconfig)

	var ptest = `
	addr = "this is addr"
	user = "xiao"
	pass = "jianhe"
	`
	type mptest struct {
		Addr string
		User string
		Pass string
	}

	var pconf mptest
	_, err := toml.Decode(ptest, &pconf)
	if err != nil {
		fmt.Println("err:%v", err)
	}
	fmt.Println(pconf)

}
