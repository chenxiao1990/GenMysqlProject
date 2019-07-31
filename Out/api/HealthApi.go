
package api
import (
	"Out/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func healthInit(groupgo *gin.RouterGroup) {
	 
	// 所有api的开头 都加上/go  例如 /go/health/create
	groupgo.POST("/health/create", HealthCreate)
	groupgo.POST("/health/delete", HealthDelete)
	groupgo.POST("/health/update", HealthUpdate)
	groupgo.POST("/health/select", HealthSelect)
}

// HealthCreate ...
func HealthCreate(c *gin.Context) {

	var param service.HealthCreate
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

// HealthDelete ...
func HealthDelete(c *gin.Context) {

	var param service.HealthDelete
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

// HealthUpdate ...
func HealthUpdate(c *gin.Context) {

	var param service.HealthUpdate
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

// HealthSelect ...
func HealthSelect(c *gin.Context) {

	var param service.HealthSelect
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

