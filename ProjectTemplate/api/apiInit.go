package api

var ApiInitTemplate = `
package api
import (
	"{{.ProjectName}}/config"
	"fmt"
	"github.com/DeanThompson/ginpprof"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"path/filepath"
	"time"
)

// GRouter 全局Router 全局调用的函数会在模块的init()之前执行
var GRouter *gin.Engine  
 
// Cors 跨域配置
func Cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Cookie"}
	// 配置允许的域名   * 代表所有都允许
	config.AllowOrigins = []string{"http://localhost:8080", "*"}
	config.AllowCredentials = true
	return cors.New(config)
}
// GinInit ...
func GinInit(){
	log.Println("启动 gin http服务 :", config.GConf.ServerPort)

	if config.GConf.OutLog {
		dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
		os.Mkdir(dir+"/log", 0777)
		f, e := os.OpenFile(dir+"/log/ginlog"+time.Now().Format("2006-01-02 15-04-05")+".log", os.O_WRONLY|os.O_CREATE|os.O_SYNC|os.O_APPEND, 0755)
		if e != nil {
			log.Println("日志文件开启失败：", e)
		} else {
			gin.DefaultWriter = f
			gin.DefaultErrorWriter = f
		}
	}

	gin.SetMode(gin.ReleaseMode)

	GRouter = gin.Default()
	ginpprof.Wrapper(GRouter)
	 
	groupgo := GRouter.Group("/go")
	// 使用跨域中间件允许跨域
	groupgo.Use(Cors())

	// 加载各个router
	initrouter(groupgo)

	
	//启动服务
	go func() {
		if err := GRouter.Run(fmt.Sprintf(":%d", config.GConf.ServerPort)); err != nil {
			log.Println(err.Error())
		}
	}()
	 
}

// 加载各个router
func initrouter(groupgo *gin.RouterGroup) {
	baseInit(groupgo)
	{{range .Inits}}{{.}}
    {{end}}
 
}
`
