package main

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/jimsmart/schema"
)

type ModelInfo struct {
	PackageName     string
	StructName      string
	ShortStructName string
	TableName       string
	Fields          []string
}

// Constants for return types of golang
const (
	golangByteArray  = "[]byte"
	gureguNullInt    = "null.Int"
	sqlNullInt       = "sql.NullInt64"
	golangInt        = "int"
	golangInt64      = "int64"
	gureguNullFloat  = "null.Float"
	sqlNullFloat     = "sql.NullFloat64"
	golangFloat      = "float"
	golangFloat32    = "float32"
	golangFloat64    = "float64"
	gureguNullString = "null.String"
	sqlNullString    = "sql.NullString"
	gureguNullTime   = "null.Time"
	golangTime       = "time.Time"
)

// GenerateStruct generates a struct for the given table.
func GenerateStruct(db *sql.DB, tableName string, structName string, pkgName string, jsonAnnotation bool, gormAnnotation bool, gureguTypes bool) *ModelInfo {
	cols, err := schema.Table(db, tableName)
	if err != nil {
		fmt.Println("查询表字段错误:", err)
		return nil
	}
	fields := generateFieldsTypes(db, tableName, cols, 0, jsonAnnotation, gormAnnotation, gureguTypes)

	//fields := generateMysqlTypes(db, columnTypes, 0, jsonAnnotation, gormAnnotation, gureguTypes)

	var modelInfo = &ModelInfo{
		PackageName:     pkgName,
		StructName:      structName,
		TableName:       tableName,
		ShortStructName: strings.ToLower(string(structName[0])),
		Fields:          fields,
	}

	return modelInfo
}

// Generate fields string
func generateFieldsTypes(db *sql.DB, tableName string, columns []*sql.ColumnType, depth int, jsonAnnotation bool, gormAnnotation bool, gureguTypes bool) []string {

	rows, err := db.Query("show create table `" + tableName + "`")

	if err != nil {
		fmt.Println("获取创建表语句出错:", err)
		return []string{}
	}
	defer rows.Close()

	createstrs := make([]string, 0)
	for rows.Next() {
		var tname string
		var cstr string
		rows.Scan(&tname, &cstr)
		tmp := strings.Split(cstr, "\n")
		createstrs = tmp[1 : len(tmp)-1]
	}

	var fields []string
	var field = ""
	for _, c := range columns {

		nullable, _ := c.Nullable()
		key := c.Name()
		dbtype := c.DatabaseTypeName()
		valueType := sqlTypeToGoType(strings.ToLower(dbtype), nullable, gureguTypes)
		if valueType == "" { // unknown type
			continue
		}
		fieldName := FmtFieldName(stringifyFirstChar(key))

		var annotations []string
		if gormAnnotation == true {
			//判断是否是主键
			zhujian := false
			zizeng := false
			suoyin := ""
			sqltype := ""

			for _, str := range createstrs {
				tmpwords := strings.Split(str, " ")
				words := make([]string, 0)
				for _, w := range tmpwords {
					if w != "" {
						w = strings.Trim(w, ",")
						words = append(words, w)
					}
				}
				// 这是一个字段
				if len(words) >= 2 && words[0] == "`"+key+"`" {
					sqltype = words[1]
				}
				//判断是不是主键
				if len(words) >= 3 && words[0] == "PRIMARY" && words[1] == "KEY" && words[2] == "(`"+key+"`)" {
					zhujian = true
				}
				//判断是不是索引
				if len(words) >= 3 && words[0] == "KEY" && words[2] == "(`"+key+"`)" {
					suoyin = "index:" + words[1][1:len(words[1])-1]
				}
				if len(words) >= 4 && words[0] == "UNIQUE" && words[1] == "KEY" && words[3] == "(`"+key+"`)" {
					suoyin = "unique_index:" + words[2][1:len(words[2])-1]
				}
				//判断是不是自增  自增不能使用type指定类型 不然自动迁移自增无效
				if words[0] == "`"+key+"`" && words[len(words)-1] == "AUTO_INCREMENT" {
					zizeng = true
				}

			}
			gormstr := fmt.Sprintf("gorm:\"column:%s", key)
			// 数据库类型
			if sqltype != "" {
				gormstr += ";type:" + sqltype
				if zizeng {
					gormstr += " AUTO_INCREMENT"
				}
			}
			//
			if zhujian {
				gormstr += ";primary_key"
			}
			if suoyin != "" {
				gormstr += fmt.Sprintf(";%s", suoyin)
			}

			gormstr += ";\""
			annotations = append(annotations, gormstr)

		}
		if jsonAnnotation == true {
			annotations = append(annotations, fmt.Sprintf("json:\"%s\"", key))
		}
		if len(annotations) > 0 {
			field = fmt.Sprintf("%s %s `%s`",
				fieldName,
				valueType,
				strings.Join(annotations, " "))

		} else {
			field = fmt.Sprintf("%s %s",
				fieldName,
				valueType)
		}

		fields = append(fields, field)
	}
	return fields
}

func sqlTypeToGoType(mysqlType string, nullable bool, gureguTypes bool) string {
	switch mysqlType {
	case "tinyint", "int", "smallint", "mediumint":

		return golangInt
	case "bigint":

		return golangInt64
	case "char", "enum", "varchar", "longtext", "mediumtext", "text", "tinytext":

		return "string"
	case "date", "datetime", "time", "timestamp":

		return golangTime
	case "decimal", "double":

		return golangFloat64
	case "float":

		return golangFloat32
	case "binary", "blob", "longblob", "mediumblob", "varbinary":
		return golangByteArray
	}
	return ""
}
