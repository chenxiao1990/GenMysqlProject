
package api
import (
	"Out/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	groupgo := GRouter.Group("/go")
	// 所有api的开头 都加上/go  例如 /go/user/create
	groupgo.POST("/user/create", UserCreate)
	groupgo.POST("/user/delete", UserDelete)
	groupgo.POST("/user/update", UserUpdate)
	groupgo.POST("/user/select", UserSelect)
}

// UserCreate ...
func UserCreate(c *gin.Context) {

	var param service.UserCreate
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

// UserDelete ...
func UserDelete(c *gin.Context) {

	var param service.UserDelete
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

// UserUpdate ...
func UserUpdate(c *gin.Context) {

	var param service.UserUpdate
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

// UserSelect ...
func UserSelect(c *gin.Context) {

	var param service.UserSelect
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

