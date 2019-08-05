

package model
import (
	"time"
)

//这里是为了有的表没有引用time而导致的编译错误
var (
	_ = time.Second
)
// Hart ...
type Hart struct {
    ID int `gorm:"column:id;primary_key;AUTO_INCREMENT;" json:"id"`
    HealHr int `gorm:"column:heal_hr;type:int(11);" json:"heal_hr"`
    Created int `gorm:"column:created;type:int(11);index:indexCreated;" json:"created"`
    Updated int `gorm:"column:updated;type:int(11);" json:"updated"`
    
}

// TableName ...
func ( Hart) TableName() string {
	return "hart"
}

 
