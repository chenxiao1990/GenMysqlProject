

package model
import (
	"time"
)

//这里是为了有的表没有引用time而导致的编译错误
var (
	_ = time.Second
)
// Health ...
type Health struct {
    ID int `gorm:"column:id;type:int(10);primary_key;AUTO_INCREMENT;" json:"id"`
    UID int `gorm:"column:uid;type:int(11);index:indexUid;" json:"uid"`
    HealStep int `gorm:"column:heal_step;type:int(11);" json:"heal_step"`
    Power int `gorm:"column:power;type:int(11);" json:"power"`
    Cal float64 `gorm:"column:cal;type:double;" json:"cal"`
    Created int `gorm:"column:created;type:int(11);" json:"created"`
    
}

// TableName ...
func ( Health) TableName() string {
	return "health"
}

 
