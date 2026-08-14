package config

import (
	"fmt"
	"github.com/tagphi/czdb-search-golang/pkg/db"
	"log"
)

var Cz88Ip *db.DBSearcher

func init() {
	InitCz88Ip()
}

// InitCz88Ip 初始化IP纯真社区版
func InitCz88Ip() {
	var err error
	Cz88Ip, err = db.InitDBSearcher("./config/cz88_public_v4.czdb", "s4s5fO8FegK89uxtvM8seg==", db.MEMORY)
	if err != nil {
		log.Fatal("初始化IP纯真社区版失败: ", err)
	}
	fmt.Println("初始化IP纯真社区版成功!")
}
