package api

var ApiTemplate = `
package api
import (
	"{{.ProjectName}}/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func {{.StructNameLow}}Init(groupgo *gin.RouterGroup) {
	 
	// 所有api的开头 都加上/go  例如 /go/{{.StructNameLow}}/create
	groupgo.POST("/{{.StructNameLow}}/create", {{.StructName}}Create)
	groupgo.POST("/{{.StructNameLow}}/delete", {{.StructName}}Delete)
	groupgo.POST("/{{.StructNameLow}}/update", {{.StructName}}Update)
	groupgo.POST("/{{.StructNameLow}}/select", {{.StructName}}Select)
}

// {{.StructName}}Create ...
func {{.StructName}}Create(c *gin.Context) {

	var param service.{{.StructName}}CreateParam
	//解析参数
	err := c.ShouldBind(&param)
	if err != nil {
		reply := NewReplyError(err.Error())
		c.JSON(http.StatusOK, reply)
		return
	}
	ser := &service.{{.StructName}}Service{}
	back, err := ser.Create(&param)
	if err != nil {
		reply := NewReplyError(err.Error())
		c.JSON(http.StatusOK, reply)
		return
	}
	reply := NewReplyOk()
	reply.Data = back
	c.JSON(http.StatusOK, reply)
	return
}

// {{.StructName}}Delete ...
func {{.StructName}}Delete(c *gin.Context) {

	var param struct {
		ID int "json:\"id\" form:\"id\""
	}
	//解析参数
	err := c.ShouldBind(&param)
	if err != nil {
		reply := NewReplyError(err.Error())
		c.JSON(http.StatusOK, reply)
		return
	}
	ser := &service.{{.StructName}}Service{}
	err = ser.Delete(param.ID)
	if err != nil {
		reply := NewReplyError(err.Error())
		c.JSON(http.StatusOK, reply)
		return
	}
	reply := NewReplyOk()
	c.JSON(http.StatusOK, reply)
	return
}

// {{.StructName}}Update ...
func {{.StructName}}Update(c *gin.Context) {

	var param service.{{.StructName}}UpdateParam
	//解析参数
	err := c.ShouldBind(&param)
	if err != nil {
		reply := NewReplyError(err.Error())
		c.JSON(http.StatusOK, reply)
		return
	}
	ser := &service.{{.StructName}}Service{}
	back, err := ser.Update(&param)
	if err != nil {
		reply := NewReplyError(err.Error())
		c.JSON(http.StatusOK, reply)
		return
	}
	reply := NewReplyOk()
	reply.Data = back
	c.JSON(http.StatusOK, reply)
	return
}

// {{.StructName}}Select ...
func {{.StructName}}Select(c *gin.Context) {

	var param struct {
		ID int "json:\"id\" form:\"id\""
	}
	//解析参数
	err := c.ShouldBindJSON(&param)
	if err != nil {
		reply := NewReplyError(err.Error())
		c.JSON(http.StatusOK, reply)
		return
	}
	ser := &service.{{.StructName}}Service{}
	back, err := ser.Select(param.ID)
	if err != nil {
		reply := NewReplyError(err.Error())
		c.JSON(http.StatusOK, reply)
		return
	}
	reply := NewReplyOk()
	reply.Data = back
	c.JSON(http.StatusOK, reply)
	return
}

`
