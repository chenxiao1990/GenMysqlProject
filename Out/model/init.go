
package model
import (
	"Out/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
)

// DB 全局数据库连接 全局调用的函数会在模块的init()之前执行
var DB *gorm.DB = initdb()

func initdb() *gorm.DB{

	// 数据连接  刚开机有可能数据未启动，所以循环直到连接成功
	for {
		log.Println("初始化数据库链接", config.GConf.Dbipport)
		if b, db := linkdb() ; b == true{

			return db
		}
		time.Sleep(1 * time.Second)
	}

}
func linkdb() (bool, *gorm.DB) {
	var err error
	dbstr := config.GConf.Dbuser + ":" + config.GConf.Dbpass + "@tcp(" + config.GConf.Dbipport + ")/" + config.GConf.Dbname + "?charset=utf8"
	tmpDB, err := gorm.Open("mysql", dbstr)
	if err != nil {
		return false,nil
	}
	//空闲
	tmpDB.DB().SetMaxIdleConns(20)
	//打开
	tmpDB.DB().SetMaxOpenConns(500)
	//超时
	tmpDB.DB().SetConnMaxLifetime(time.Second * 60)
	return true, tmpDB
}

