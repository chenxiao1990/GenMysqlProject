

package model
import (
	"time"
)

//这里是为了有的表没有引用time而导致的编译错误
var (
	_ = time.Second
)
// Check ...
type Check struct {
    ID int `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT;" json:"id"`
    UID int `gorm:"column:uid;type:int(11);index:indexUid;" json:"uid"`
    AreaID string `gorm:"column:area_id;type:varchar(20);" json:"area_id"`
    GradeID int `gorm:"column:grade_id;type:int(11);index:indexGid;" json:"grade_id"`
    ClassID int `gorm:"column:class_id;type:int(11);index:indexCid;" json:"class_id"`
    Created int `gorm:"column:created;type:int(11);" json:"created"`
    Updated int `gorm:"column:updated;type:int(11);" json:"updated"`
    
}

// TableName ...
func ( Check) TableName() string {
	return "check"
}

 
