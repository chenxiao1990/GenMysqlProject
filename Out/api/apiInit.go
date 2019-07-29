package api
  
import (
	"Out/config"
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
var GRouter *gin.Engine = initgin()
 
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
func initgin() *gin.Engine{
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

	tmpGRouter := gin.Default()
	ginpprof.Wrapper(tmpGRouter)
	 
 	// 使用跨域中间件允许跨域
	tmpGRouter.Use(Cors())
	
	//启动服务
	go func() {
		if err := tmpGRouter.Run(fmt.Sprintf(":%d", config.GConf.ServerPort)); err != nil {
			log.Println(err.Error())
		}
	}()
	return tmpGRouter
}
