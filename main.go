package main

import (
	"GenMysqlProject/ProjectTemplate/config"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/droundy/goopt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jimsmart/schema"
	"github.com/jinzhu/inflection"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var (
	dbIPPort       = goopt.String([]string{"--ipport", "-i"}, "192.168.0.86:3306", "database ipport")
	dbName         = goopt.String([]string{"--dbname", "-d"}, "cx", "Database name")
	dbUser         = goopt.String([]string{"--user", "-u"}, "root", "Database user name")
	dbPass         = goopt.String([]string{"--pass", "-p"}, "12345678", "Database user pass")
	outProjectName = goopt.String([]string{"--outname", "-o"}, "Out", "name to set for out Project")
)

func init() {

	// Setup goopts
	goopt.Description = func() string {
		return "GORM GIN API generator for Mysql"
	}
	goopt.Version = "0.0.1"
	goopt.Summary = `.\GenMysqlProject.exe --ipport "192.168.0.86:3306" --dbname "cx" --user "root" --pass "12345678" --outname "mypro" `

	//Parse options
	goopt.Parse(nil)

}

type Base struct {
	ProjectName string
}

func main() {
	os.Mkdir(*outProjectName, 0777)
	os.Mkdir(*outProjectName+"/api", 0777)
	os.Mkdir(*outProjectName+"/config", 0777)
	os.Mkdir(*outProjectName+"/dao", 0777)
	os.Mkdir(*outProjectName+"/model", 0777)
	os.Mkdir(*outProjectName+"/service", 0777)

	curdir, _ := filepath.Abs("./")

	base := Base{
		ProjectName: *outProjectName,
	}

	fromdir := curdir + "/ProjectTemplate"
	todir := curdir + "/" + *outProjectName

	// main文件
	basefile(fromdir+"/main.go", todir+"/main.go", base)
	// dockerfile
	basefile(fromdir+"/Dockerfile", todir+"/Dockerfile", base)
	// config
	copyfile(fromdir+"/config/config.go", todir+"/config/config.go")
	config.GConf.ServerPort = 80
	config.GConf.OutLog = false
	config.GConf.Dbipport = *dbIPPort
	config.GConf.Dbuser = *dbUser
	config.GConf.Dbpass = *dbPass
	config.GConf.Dbname = *dbName

	configbb, _ := json.Marshal(config.GConf)
	os.MkdirAll(todir+"/config", 0666)
	configf, _ := os.OpenFile(todir+"/config/config.json", os.O_WRONLY|os.O_CREATE|os.O_SYNC, 0755)
	configf.Write(configbb)
	configf.Close()

	// 搞model文件夹内的
	var tables = make([]string, 0)
	{
		basefile(fromdir+"/model/init.go", todir+"/model/init.go", base)
		// 获取所有数据库表
		dbstr := *dbUser + ":" + *dbPass + "@tcp(" + *dbIPPort + ")/" + *dbName + "?charset=utf8"
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

		// generate go files for each table
		for _, tableName := range tables {
			structName := FmtFieldName(tableName)
			structName = inflection.Singular(structName)

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
			basefile(fromdir+"/model/table.go", todir+"/model/"+tableName+"Model.go", base)
		}
	}

	// 搞service文件夹内的
	{
		dbstr := *dbUser + ":" + *dbPass + "@tcp(" + *dbIPPort + ")/" + *dbName + "?charset=utf8"
		var db, err = sql.Open("mysql", dbstr)
		if err != nil {
			fmt.Println("Error in open database: " + err.Error())
			return
		}
		defer db.Close()
		// generate go files for each table
		for _, tableName := range tables {
			structName := FmtFieldName(tableName)
			structName = inflection.Singular(structName)

			modelInfo := GenerateStruct(db, tableName, structName, "model", true, false, true)

			var base = struct {
				ProjectName  string
				StructName   string
				FieldsCreate []string
				Fields       []string
			}{
				ProjectName:  *outProjectName,
				StructName:   structName,
				FieldsCreate: modelInfo.Fields, //这里应该去掉主键，但是懒得弄了
				Fields:       modelInfo.Fields,
			}
			basefile(fromdir+"/service/service.go", todir+"/service/"+tableName+"Service.go", base)
		}
	}

	// 搞dao文件夹内的
	{
		dbstr := *dbUser + ":" + *dbPass + "@tcp(" + *dbIPPort + ")/" + *dbName + "?charset=utf8"
		var db, err = sql.Open("mysql", dbstr)
		if err != nil {
			fmt.Println("Error in open database: " + err.Error())
			return
		}
		defer db.Close()
		// generate go files for each table
		for _, tableName := range tables {
			structName := FmtFieldName(tableName)
			structName = inflection.Singular(structName)

			var base = struct {
				ProjectName string
				StructName  string
			}{
				ProjectName: *outProjectName,
				StructName:  structName,
			}
			basefile(fromdir+"/dao/dao.go", todir+"/dao/"+tableName+"Dao.go", base)
		}
	}
	// 搞api文件夹内的
	{
		basefile(fromdir+"/api/apiInit.go", todir+"/api/apiInit.go", base)
		copyfile(fromdir+"/api/apiReply.go", todir+"/api/apiReply.go")
		copyfile(fromdir+"/api/baseApi.go", todir+"/api/baseApi.go")

		dbstr := *dbUser + ":" + *dbPass + "@tcp(" + *dbIPPort + ")/" + *dbName + "?charset=utf8"
		var db, err = sql.Open("mysql", dbstr)
		if err != nil {
			fmt.Println("Error in open database: " + err.Error())
			return
		}
		defer db.Close()
		// generate go files for each table
		for _, tableName := range tables {
			structName := FmtFieldName(tableName)
			structName = inflection.Singular(structName)

			var base = struct {
				ProjectName   string
				StructNameLow string
				StructName    string
			}{
				ProjectName:   *outProjectName,
				StructNameLow: strings.ToLower(structName),
				StructName:    structName,
			}
			basefile(fromdir+"/api/api.go", todir+"/api/"+tableName+"Api.go", base)
		}
	}
}

// copyfile
func copyfile(frompath string, topath string) {

	bb, _ := ioutil.ReadFile(frompath)

	todir := filepath.Dir(topath)
	os.MkdirAll(todir, 0666)
	mainf, _ := os.OpenFile(topath, os.O_WRONLY|os.O_CREATE|os.O_SYNC, 0755)
	mainf.Write(bb)
	mainf.Close()
}

// 基础文件拷贝
func basefile(frompath string, topath string, base interface{}) {

	tmpl, err := template.ParseFiles(frompath) //建立一个模板
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
