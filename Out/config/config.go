
package config
import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Config 配置文件
type Config struct {
	ServerPort int    // api服务端口
	OutLog     bool   //是否输出日志文件
	Dbipport   string //数据ip:port
	Dbuser     string //用户名
	Dbpass     string //密码
	Dbname     string //数据库名
}

// GConf 全局的config变量
var GConf Config


// Loadconfig ...
func Loadconfig(configfile string) {
	 
	btMsg, _ := ioutil.ReadFile(configfile)
	//utf8去掉bom的头
	btMsg = bytes.TrimPrefix(btMsg, []byte("\xef\xbb\xbf"))
	if err := json.Unmarshal(btMsg, &GConf); err != nil {
		panic(0)
	} else {
		fmt.Println("读取配置文件成功:", fmt.Sprintf("%+v", GConf))
	}
}

