
package api
import (
	"Out/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func hartInit(groupgo *gin.RouterGroup) {
	 
	// 所有api的开头 都加上/go  例如 /go/hart/create
	groupgo.POST("/hart/create", HartCreate)
	groupgo.POST("/hart/delete", HartDelete)
	groupgo.POST("/hart/update", HartUpdate)
	groupgo.POST("/hart/select", HartSelect)
}

// HartCreate ...
func HartCreate(c *gin.Context) {

	var param service.HartCreate
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

// HartDelete ...
func HartDelete(c *gin.Context) {

	var param service.HartDelete
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

// HartUpdate ...
func HartUpdate(c *gin.Context) {

	var param service.HartUpdate
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

// HartSelect ...
func HartSelect(c *gin.Context) {

	var param service.HartSelect
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

