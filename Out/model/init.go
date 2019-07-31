
package model
import (
	"Out/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
)

// DB 全局数据库连接
var DB *gorm.DB

// LinkDB 初始化连接db
func LinkDB() {

	// 数据连接  刚开机有可能数据未启动，所以循环直到连接成功
	for {
		log.Println("初始化数据库链接", config.GConf.Dbipport)
		if b, db := link(); b == true {
			DB = db
			AutoMigrate()
			return 
		}
		time.Sleep(1 * time.Second)
	}

}
func link() (bool, *gorm.DB) {
	var err error
	dbstr := config.GConf.Dbuser + ":" + config.GConf.Dbpass + "@tcp(" + config.GConf.Dbipport + ")/" + config.GConf.Dbname + "?charset=utf8"
	tmpDB, err := gorm.Open("mysql", dbstr)
	if err != nil {
		return false, nil
	}
	//空闲
	tmpDB.DB().SetMaxIdleConns(20)
	//打开
	tmpDB.DB().SetMaxOpenConns(500)
	//超时
	tmpDB.DB().SetConnMaxLifetime(time.Second * 60)
	return true, tmpDB
}

// AutoMigrate ...
func AutoMigrate() {
	
	if er := DB.AutoMigrate(&Check{}).Error; er != nil {
		log.Println("自动迁移错误:", er)
	}
    
	if er := DB.AutoMigrate(&Hart{}).Error; er != nil {
		log.Println("自动迁移错误:", er)
	}
    
	if er := DB.AutoMigrate(&Health{}).Error; er != nil {
		log.Println("自动迁移错误:", er)
	}
    
}

