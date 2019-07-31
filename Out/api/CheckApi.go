
package api
import (
	"Out/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func checkInit(groupgo *gin.RouterGroup) {
	 
	// 所有api的开头 都加上/go  例如 /go/check/create
	groupgo.POST("/check/create", CheckCreate)
	groupgo.POST("/check/delete", CheckDelete)
	groupgo.POST("/check/update", CheckUpdate)
	groupgo.POST("/check/select", CheckSelect)
}

// CheckCreate ...
func CheckCreate(c *gin.Context) {

	var param service.CheckCreate
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

// CheckDelete ...
func CheckDelete(c *gin.Context) {

	var param service.CheckDelete
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

// CheckUpdate ...
func CheckUpdate(c *gin.Context) {

	var param service.CheckUpdate
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

// CheckSelect ...
func CheckSelect(c *gin.Context) {

	var param service.CheckSelect
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

