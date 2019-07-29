package api

import (
	"{{.ProjectName}}/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	groupgo := GRouter.Group("/go")
	// 所有api的开头 都加上/go  例如 /go/{{.StructNameLow}}/create
	groupgo.POST("/{{.StructNameLow}}/create", {{.StructName}}Create)
	groupgo.POST("/{{.StructNameLow}}/delete", {{.StructName}}Delete)
	groupgo.POST("/{{.StructNameLow}}/update", {{.StructName}}Update)
	groupgo.POST("/{{.StructNameLow}}/select", {{.StructName}}Select)
}

// {{.StructName}}Create ...
func {{.StructName}}Create(c *gin.Context) {

	var param service.{{.StructName}}Create
	//解析参数
	err := c.ShouldBindJSON(&param)
	if err != nil {
		reply := NewReplyError(err.Error())
		c.JSON(http.StatusOK, reply)
		return
	}

	back, err := param.Create()
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

	var param service.{{.StructName}}Delete
	//解析参数
	err := c.ShouldBindJSON(&param)
	if err != nil {
		reply := NewReplyError(err.Error())
		c.JSON(http.StatusOK, reply)
		return
	}

	err = param.Delete()
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

	var param service.{{.StructName}}Update
	//解析参数
	err := c.ShouldBindJSON(&param)
	if err != nil {
		reply := NewReplyError(err.Error())
		c.JSON(http.StatusOK, reply)
		return
	}

	back, err := param.Update()
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

	var param service.{{.StructName}}Select
	//解析参数
	err := c.ShouldBindJSON(&param)
	if err != nil {
		reply := NewReplyError(err.Error())
		c.JSON(http.StatusOK, reply)
		return
	}

	back, err := param.Select()
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
