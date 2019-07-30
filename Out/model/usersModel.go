

package model
import (
	"log"
	"time"
)

//这里是为了有的表没有引用time而导致的编译错误
var (
	_ = time.Second
)
// User ...
type User struct {
    ID int `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT;" json:"id"`
    UID int `gorm:"column:uid;type:int(11);" json:"uid"`
    Name string `gorm:"column:name;type:varchar(255);index:idx_name_code;" json:"name"`
    Phone string `gorm:"column:phone;type:text;" json:"phone"`
    Pass string `gorm:"column:pass;type:varchar(255);" json:"pass"`
    Status int `gorm:"column:status;type:int(11);" json:"status"`
    
}

// TableName ...
func ( User) TableName() string {
	return "users"
}


 
// 自动迁移
func init() {
	if er := DB.AutoMigrate(&User{}).Error; er != nil {
		log.Println("自动迁移错误:", er)
	}
}

