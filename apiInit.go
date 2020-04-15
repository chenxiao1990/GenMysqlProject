package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
	"unsafe"

	"github.com/chenxiao1990/GenMysqlProject/ProjectTemplate/model"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jimsmart/schema"
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

	fuc := cors.New(config)
	up := unsafe.Pointer(&fuc)
	return *((*gin.HandlerFunc)(up))
}

// GinInit ...
func GinInit(port int) {
	log.Println("启动 gin http服务 :", port)

	var param struct {
		DbIPPort       string `json:"dbIPPort" binding:"required"`
		DbName         string `json:"dbName" binding:"required"`
		DbUser         string `json:"dbUser" binding:"required"`
		DbPass         string `json:"dbPass" binding:"required"`
		OutProjectName string `json:"outProjectName" binding:"required"`
	}

	//文件
	bb, _ := ioutil.ReadFile("dbinfo.json")
	json.Unmarshal(bb, &param)
	if param.DbIPPort != "" &&
		param.DbName != "" &&
		param.DbUser != "" &&
		param.DbPass != "" {
		dbIPPort = param.DbIPPort
		dbName = param.DbName
		dbUser = param.DbUser
		dbPass = param.DbPass
		outProjectName = param.OutProjectName
	}

	gin.SetMode(gin.ReleaseMode)

	GRouter = gin.Default()
	// 使用权限检测中间件 使用跨域中间件允许跨域
	GRouter.Use(Cors())

	groupgo := GRouter.Group("/cx")

	// 加载各个router
	initrouter(groupgo)

	//启动服务
	go func() {
		if err := GRouter.Run(fmt.Sprintf(":%d", port)); err != nil {
			log.Println(err.Error())
		}
	}()

}

// Reply api的回复结构
type Reply struct {
	Code    int         "json:\"code\""
	Message string      "json:\"message\""
	Data    interface{} "json:\"data\""
}

// NewReplyOk 默认的正确回复
func NewReplyOk() *Reply {
	return &Reply{
		Code:    1,
		Message: "成功",
		Data:    make(map[string]interface{}),
	}
}

// NewReplyError 默认的错误回复
func NewReplyError(msg string) *Reply {
	return &Reply{
		Code:    0,
		Message: msg,
		Data:    make(map[string]interface{}),
	}
}

// 加载各个router
func initrouter(groupgo *gin.RouterGroup) {
	// 静态文件页面
	//dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	//GRouter.Static("/vue", dir+"/static/vue")

	GRouter.StaticFS("/vue", AssetFile())

	groupgo.GET("/version", func(c *gin.Context) {
		reply := NewReplyOk()
		reply.Data = "1.2.2"
		c.JSON(http.StatusOK, reply)
	})

	groupgo.POST("/genproject", genproject)
	groupgo.POST("/gentablestruct", gentablestruct)

	groupgo.GET("/dbinfo", dbinfo)
	groupgo.POST("/setdbinfo", setdbinfo)

	groupgo.GET("/tables", tables)

	groupgo.POST("/format", codeformat)
}

func codeformat(c *gin.Context) {
	var param struct {
		Code string `json:"code" binding:"required"`
	}
	err := c.ShouldBind(&param)
	if err != nil {
		reply := NewReplyError(err.Error())
		c.JSON(http.StatusOK, reply)
		return
	}
	bb, err := format.Source([]byte(param.Code))
	reply := NewReplyOk()
	reply.Data = string(bb)
	c.JSON(http.StatusOK, reply)
}
func tables(c *gin.Context) {
	// 获取所有数据库表
	dbstr := dbUser + ":" + dbPass + "@tcp(" + dbIPPort + ")/" + dbName + "?charset=utf8&parseTime=true&loc=Local"
	var db, err = sql.Open("mysql", dbstr)
	if err != nil {
		fmt.Println("Error in open database: " + err.Error())
		return
	}
	defer db.Close()

	tables, err := schema.TableNames(db)
	if err != nil {
		reply := NewReplyError(err.Error())
		c.JSON(http.StatusOK, reply)
		return
	}

	back := make([]interface{}, 0)
	for _, tableName := range tables {
		structName := FmtFieldName(tableName)
		if structName[len(structName)-1] == 's' {
			structName = structName[0 : len(structName)-1]
		}
		cols, _ := schema.Table(db, tableName)

		type Filed struct {
			Name string
			Type string
		}
		fileds := make([]Filed, 0)
		for _, col := range cols {
			f := Filed{
				Name: col.Name(),
				Type: col.DatabaseTypeName(),
			}
			fileds = append(fileds, f)
		}
		var base = struct {
			TableName  string
			StructName string
			Fields     []Filed
		}{
			TableName:  tableName,
			StructName: structName,
			Fields:     fileds,
		}
		back = append(back, base)
	}
	reply := NewReplyOk()
	reply.Data = back
	c.JSON(http.StatusOK, reply)
}
func setdbinfo(c *gin.Context) {
	var param struct {
		DbIPPort string `json:"dbIPPort" binding:"required"`
		DbName   string `json:"dbName" binding:"required"`
		DbUser   string `json:"dbUser" binding:"required"`
		DbPass   string `json:"dbPass" binding:"required"`
	}
	err := c.ShouldBind(&param)
	if err != nil {
		reply := NewReplyError(err.Error())
		c.JSON(http.StatusOK, reply)
		return
	}

	//把配置保存到文件
	bb, _ := json.Marshal(param)
	ioutil.WriteFile("dbinfo.json", bb, os.ModePerm)

	dbIPPort = param.DbIPPort
	dbName = param.DbName
	dbUser = param.DbUser
	dbPass = param.DbPass

	reply := NewReplyOk()
	c.JSON(http.StatusOK, reply)
}
func dbinfo(c *gin.Context) {
	var param = struct {
		DbIPPort       string `json:"dbIPPort" binding:"required"`
		DbName         string `json:"dbName" binding:"required"`
		DbUser         string `json:"dbUser" binding:"required"`
		DbPass         string `json:"dbPass" binding:"required"`
		OutProjectName string `json:"outProjectName" binding:"required"`
	}{
		DbIPPort:       dbIPPort,
		DbName:         dbName,
		DbUser:         dbUser,
		DbPass:         dbPass,
		OutProjectName: outProjectName,
	}

	reply := NewReplyOk()
	reply.Data = param
	c.JSON(http.StatusOK, reply)
}
func gentablestruct(c *gin.Context) {
	type Param struct {
		StructName string
		TableName  string
	}
	var param Param
	err := c.ShouldBindJSON(&param)
	if err != nil {
		reply := NewReplyError(err.Error())
		c.JSON(http.StatusOK, reply)
		return
	}
	// 获取所有数据库表
	dbstr := dbUser + ":" + dbPass + "@tcp(" + dbIPPort + ")/" + dbName + "?charset=utf8&parseTime=true&loc=Local"
	db, err := sql.Open("mysql", dbstr)
	if err != nil {

		reply := NewReplyError("Error in open database: " + err.Error())
		c.JSON(http.StatusOK, reply)
		return
	}
	defer db.Close()
	modelInfo := GenerateStruct(db, param.TableName, param.StructName, "model", true, true, true)
	var base = struct {
		StructName string
		TableName  string
		Fields     []string
	}{
		StructName: param.StructName,
		TableName:  param.TableName,
		Fields:     modelInfo.Fields,
	}
	tmpl, err := template.New("base").Parse(model.TableTemplate) //建立一个模板
	if err != nil {
		reply := NewReplyError(err.Error())
		c.JSON(http.StatusOK, reply)
		return
	}

	out := bytes.NewBuffer([]byte{})
	err = tmpl.Execute(out, base)
	if err != nil {
		reply := NewReplyError(err.Error())
		c.JSON(http.StatusOK, reply)
		return
	}
	bb, err := format.Source(out.Bytes())
	if err != nil {
		reply := NewReplyError(err.Error())
		c.JSON(http.StatusOK, reply)
		return
	}
	reply := NewReplyOk()
	reply.Data = string(bb)
	c.JSON(http.StatusOK, reply)

}
func genproject(c *gin.Context) {
	var param struct {
		DbIPPort       string `json:"dbIPPort" binding:"required"`
		DbName         string `json:"dbName" binding:"required"`
		DbUser         string `json:"dbUser" binding:"required"`
		DbPass         string `json:"dbPass" binding:"required"`
		OutProjectName string `json:"outProjectName" binding:"required"`
	}
	err := c.ShouldBind(&param)
	if err != nil {
		reply := NewReplyError(err.Error())
		c.JSON(http.StatusOK, reply)
		return
	}

	//把配置保存到文件
	bb, _ := json.Marshal(param)
	ioutil.WriteFile("dbinfo.json", bb, os.ModePerm)

	dbIPPort = param.DbIPPort
	dbName = param.DbName
	dbUser = param.DbUser
	dbPass = param.DbPass
	outProjectName = param.OutProjectName

	curdir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	prodir := filepath.Join(curdir, outProjectName)
	_, err = os.Stat(prodir)
	if err == nil {
		//已存在
		reply := NewReplyError("当前位置已存在此目录" + outProjectName)
		c.JSON(http.StatusOK, reply)
		return
	}

	GenProject()
	reply := NewReplyOk()
	c.JSON(http.StatusOK, reply)
}
