package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/chenxiao1990/GenMysqlProject/ProjectTemplate"
	"github.com/chenxiao1990/GenMysqlProject/ProjectTemplate/api"
	"github.com/chenxiao1990/GenMysqlProject/ProjectTemplate/config"
	"github.com/chenxiao1990/GenMysqlProject/ProjectTemplate/dao"
	"github.com/chenxiao1990/GenMysqlProject/ProjectTemplate/model"
	"github.com/chenxiao1990/GenMysqlProject/ProjectTemplate/service"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jimsmart/schema"
)

var (
	dbIPPort       string
	dbName         string
	dbUser         string
	dbPass         string
	outProjectName string
)

type Base struct {
	ProjectName string
}

func GenProject() {

	os.Mkdir(outProjectName, 0777)
	os.Mkdir(outProjectName+"/api", 0777)
	os.Mkdir(outProjectName+"/config", 0777)
	os.Mkdir(outProjectName+"/dao", 0777)
	os.Mkdir(outProjectName+"/model", 0777)
	os.Mkdir(outProjectName+"/service", 0777)

	curdir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	base := Base{
		ProjectName: outProjectName,
	}

	todir := curdir + "/" + outProjectName

	// main文件
	basestr(ProjectTemplate.MainTemplate, todir+"/main.go", base)

	// config
	copytofile(config.Configstr, todir+"/config/config.go")
	config.GConf.ServerPort = 80
	config.GConf.OutLog = false
	config.GConf.Dbipport = dbIPPort
	config.GConf.Dbuser = dbUser
	config.GConf.Dbpass = dbPass
	config.GConf.Dbname = dbName

	configbb, _ := json.Marshal(config.GConf)
	os.MkdirAll(todir+"/config", 0666)
	configf, _ := os.OpenFile(todir+"/config/config.json", os.O_WRONLY|os.O_CREATE|os.O_SYNC, 0755)
	configf.Write(configbb)
	configf.Close()

	// go mod
	{
		modf, _ := os.OpenFile(todir+"/go.mod", os.O_WRONLY|os.O_CREATE|os.O_SYNC, 0755)
		modf.Write([]byte(`module ` + outProjectName + `

go 1.14`))
		modf.Close()
	}

	// 搞model文件夹内的
	var tables = make([]string, 0)
	{

		// 获取所有数据库表
		dbstr := dbUser + ":" + dbPass + "@tcp(" + dbIPPort + ")/" + dbName + "?charset=utf8&parseTime=true&loc=Local"
		var db, err = sql.Open("mysql", dbstr)
		if err != nil {
			fmt.Println("Error in open database: " + err.Error())
			return
		}
		defer db.Close()

		// parse or read tables
		tables, err = schema.TableNames(db)
		if err != nil {
			fmt.Println("Error in fetching tables information from mysql information schema", err)
			return
		}
		automigrate := make([]string, 0)
		// generate go files for each table
		for _, tableName := range tables {
			structName := FmtFieldName(tableName)
			if structName[len(structName)-1] == 's' {
				structName = structName[0 : len(structName)-1]
			}

			autostr := `
	if er := DB.AutoMigrate(&` + structName + `{}).Error; er != nil {
		log.Println("自动迁移错误:", er)
	}`
			automigrate = append(automigrate, autostr)

			modelInfo := GenerateStruct(db, tableName, structName, "model", true, true, true)

			var base = struct {
				StructName string
				TableName  string
				Fields     []string
			}{
				StructName: structName,
				TableName:  tableName,
				Fields:     modelInfo.Fields,
			}
			basestr(model.TableTemplate, todir+"/model/"+tableName+"Model.go", base)
		}
		// 搞一下 model下的init文件
		{
			var base = struct {
				ProjectName string
				AutoMigrate []string
			}{
				ProjectName: outProjectName,
				AutoMigrate: automigrate,
			}
			basestr(model.InitTemplate, todir+"/model/init.go", base)
		}

	}

	// 搞service文件夹内的
	{
		dbstr := dbUser + ":" + dbPass + "@tcp(" + dbIPPort + ")/" + dbName + "?charset=utf8"
		var db, err = sql.Open("mysql", dbstr)
		if err != nil {
			fmt.Println("Error in open database: " + err.Error())
			return
		}
		defer db.Close()
		// generate go files for each table
		for _, tableName := range tables {
			structName := FmtFieldName(tableName)
			if structName[len(structName)-1] == 's' {
				structName = structName[0 : len(structName)-1]
			}

			modelInfo := GenerateStruct(db, tableName, structName, "model", true, false, true)

			var base = struct {
				ProjectName  string
				StructName   string
				FieldsCreate []string
				Fields       []string
			}{
				ProjectName:  outProjectName,
				StructName:   structName,
				FieldsCreate: modelInfo.Fields, //这里应该去掉主键，但是懒得弄了
				Fields:       modelInfo.Fields,
			}
			basestr(service.ServiceTemplate, todir+"/service/"+tableName+"Service.go", base)
		}
	}

	// 搞dao文件夹内的
	{
		dbstr := dbUser + ":" + dbPass + "@tcp(" + dbIPPort + ")/" + dbName + "?charset=utf8"
		var db, err = sql.Open("mysql", dbstr)
		if err != nil {
			fmt.Println("Error in open database: " + err.Error())
			return
		}
		defer db.Close()
		// generate go files for each table
		for _, tableName := range tables {
			structName := FmtFieldName(tableName)
			if structName[len(structName)-1] == 's' {
				structName = structName[0 : len(structName)-1]
			}

			var base = struct {
				ProjectName string
				StructName  string
			}{
				ProjectName: outProjectName,
				StructName:  structName,
			}
			basestr(dao.DapTemplate, todir+"/dao/"+tableName+"Dao.go", base)
		}
	}
	// 搞api文件夹内的
	{

		copytofile(api.ApiReplystr, todir+"/api/apiReply.go")
		copytofile(api.ApiBasestr, todir+"/api/baseApi.go")

		dbstr := dbUser + ":" + dbPass + "@tcp(" + dbIPPort + ")/" + dbName + "?charset=utf8"
		var db, err = sql.Open("mysql", dbstr)
		if err != nil {
			fmt.Println("Error in open database: " + err.Error())
			return
		}
		defer db.Close()
		// generate go files for each table
		var initfuncs = make([]string, 0)
		for _, tableName := range tables {
			structName := FmtFieldName(tableName)
			if structName[len(structName)-1] == 's' {
				structName = structName[0 : len(structName)-1]
			}

			var base = struct {
				ProjectName   string
				StructNameLow string
				StructName    string
			}{
				ProjectName:   outProjectName,
				StructNameLow: strings.ToLower(structName),
				StructName:    structName,
			}
			basestr(api.ApiTemplate, todir+"/api/"+tableName+"Api.go", base)
			initfuncs = append(initfuncs, fmt.Sprintf("%sInit(groupgo)", strings.ToLower(structName)))
		}

		var apiinit = struct {
			ProjectName string
			Inits       []string
		}{
			ProjectName: outProjectName,
			Inits:       initfuncs,
		}
		basestr(api.ApiInitTemplate, todir+"/api/apiInit.go", apiinit)
	}
}

// copyfile
func copytofile(filestr string, topath string) {

	bb := []byte(filestr)

	todir := filepath.Dir(topath)
	os.MkdirAll(todir, 0666)
	mainf, _ := os.OpenFile(topath, os.O_WRONLY|os.O_CREATE|os.O_SYNC, 0755)
	mainf.Write(bb)
	mainf.Close()
}

// 基础文件拷贝
func basestr(basestr string, topath string, base interface{}) {

	tmpl, err := template.New("base").Parse(basestr) //建立一个模板
	if err != nil {
		fmt.Println(err)
		return
	}
	type Main struct {
		ProjectName string
	}
	todir := filepath.Dir(topath)
	os.MkdirAll(todir, 0666)
	mainf, _ := os.OpenFile(topath, os.O_WRONLY|os.O_CREATE|os.O_SYNC, 0755)
	err = tmpl.Execute(mainf, base)
	if err != nil {
		fmt.Println(err)
		return
	}
	mainf.Close()
}
