
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

	var param service.HartCreateParam
	//解析参数
	err := c.ShouldBindJSON(&param)
	if err != nil {
		reply := NewReplyError(err.Error())
		c.JSON(http.StatusOK, reply)
		return
	}
	ser := &service.HartService{}
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

// HartDelete ...
func HartDelete(c *gin.Context) {

	var param service.HartDeleteParam
	//解析参数
	err := c.ShouldBindJSON(&param)
	if err != nil {
		reply := NewReplyError(err.Error())
		c.JSON(http.StatusOK, reply)
		return
	}
	ser := &service.HartService{}
	err = ser.Delete(&param)
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

	var param service.HartUpdateParam
	//解析参数
	err := c.ShouldBindJSON(&param)
	if err != nil {
		reply := NewReplyError(err.Error())
		c.JSON(http.StatusOK, reply)
		return
	}
	ser := &service.HartService{}
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

// HartSelect ...
func HartSelect(c *gin.Context) {

	var param service.HartSelectParam
	//解析参数
	err := c.ShouldBindJSON(&param)
	if err != nil {
		reply := NewReplyError(err.Error())
		c.JSON(http.StatusOK, reply)
		return
	}
	ser := &service.HartService{}
	back, err := ser.Select(&param)
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

