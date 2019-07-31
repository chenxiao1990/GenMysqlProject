package ProjectTemplate

var MainTemplate = `
package main

import (
	"{{.ProjectName}}/api"
	"{{.ProjectName}}/config"
	_ "{{.ProjectName}}/dao"
	"{{.ProjectName}}/model"
	_ "{{.ProjectName}}/service"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {

	//加载配置文件
	var GCurDir, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	configfile := GCurDir + "/config/config.json"
	config.Loadconfig(configfile)

	// 初始化路由
	api.GinInit()
	//连接数据库
	model.LinkDB()

	// 定义输出日志文件
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	if config.GConf.OutLog {
		dir, _ := filepath.Abs("./")
		os.Mkdir(dir+"/log", 0777)
		go func() {
			for {
				f, e := os.OpenFile(dir+"/log/log"+time.Now().Format("2006-01-02 15-04-05")+".log", os.O_WRONLY|os.O_CREATE|os.O_SYNC|os.O_APPEND, 0755)
				if e != nil {
					log.Println("日志文件开启失败：", e)
					time.Sleep(1 * time.Second)
				} else {
					log.SetOutput(f)
					log.Println("-------------------")
					time.Sleep(24 * time.Hour)
					f.Close()
				}

			}

		}()
	}

	select {}
}

`
