package main

import (
	_ "{{.ProjectName}}/api"
	"{{.ProjectName}}/config"
	_ "{{.ProjectName}}/dao"
	_ "{{.ProjectName}}/model"
	_ "{{.ProjectName}}/service"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {

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
