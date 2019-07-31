
package api
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 加载路由
func baseInit(groupgo *gin.RouterGroup) {
	  
	// 所有api的开头 都加上/go  例如 /go/ping
	groupgo.GET("ping", Ping)

}

// Ping 测试存活
func Ping(c *gin.Context) {
	reply := NewReplyOk()
	c.JSON(http.StatusOK, reply)
}

