package model

import (
	"log"
	"time"
)

//这里是为了有的表没有引用time而导致的编译错误
var (
	_ = time.Second
)
// {{.StructName}} ...
type {{.StructName}} struct {
    {{range .Fields}}{{.}}
    {{end}}
}

// TableName ...
func ( {{.StructName}}) TableName() string {
	return "{{.TableName}}"
}


 
// 自动迁移
func init() {
	if er := DB.AutoMigrate(&{{.StructName}}{}).Error; er != nil {
		log.Println("自动迁移错误:", er)
	}
}
