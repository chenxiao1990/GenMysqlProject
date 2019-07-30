package service

var ServiceTemplate = `

package service
import (
	"{{.ProjectName}}/dao"
	"{{.ProjectName}}/model"
	"encoding/json"
)

// {{.StructName}}Create  服务：创建  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type {{.StructName}}Create struct {
	{{range .FieldsCreate}}{{.}}
    {{end}}
}

// {{.StructName}}CreateBack  返回参数
type {{.StructName}}CreateBack struct {
	{{range .Fields}}{{.}}
    {{end}}
}

// Create 创建
func (p *{{.StructName}}Create) Create() (*{{.StructName}}CreateBack, error) {
	dao := &dao.{{.StructName}}Dao{}
	var model  model.{{.StructName}} 
	//这里是模板创建生成的，正常应该输入参数根据实际情况赋值给dao需要的参数
	tmpbb , _ :=  json.Marshal(*p)
	json.Unmarshal(tmpbb, &model)
	data, err := dao.Create{{.StructName}}(&model)

	if err != nil {
		return nil, err

	}

	var back {{.StructName}}CreateBack
	//这里是模板创建生成的，正常应该输入参数根据实际情况赋值给dao需要的参数
	tmpbb , _  =  json.Marshal(*data)
	json.Unmarshal(tmpbb, &back)
 
	return &back, nil
}

// {{.StructName}}Delete  服务：删除   
type {{.StructName}}Delete struct {
	ID int "json:\"id\" binding:\"required\""
}

// Delete 创建
func (p *{{.StructName}}Delete) Delete() error {
	dao := &dao.{{.StructName}}Dao{}
	return dao.Delete{{.StructName}}(p.ID)
}

// {{.StructName}}Select  服务：查询
type {{.StructName}}Select struct {
	ID int "json:\"id\" binding:\"required\""
}

// {{.StructName}}SelectBack  返回参数
type {{.StructName}}SelectBack struct {
	{{range .Fields}}{{.}}
    {{end}}
}

// Select ...
func (p *{{.StructName}}Select) Select() (*{{.StructName}}SelectBack, error) {
	dao := &dao.{{.StructName}}Dao{}
	data, err := dao.Select{{.StructName}}ByID(p.ID)
	if err != nil {
		return nil, err
	}
	var back {{.StructName}}SelectBack 
	//这里是模板创建生成的，正常应该输入参数根据实际情况赋值给dao需要的参数
	tmpbb , _  :=  json.Marshal(*data)
	json.Unmarshal(tmpbb, &back)

	return &back, nil
}

// {{.StructName}}Update  服务：更新
type {{.StructName}}Update struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// {{.StructName}}UpdateBack  返回参数
type {{.StructName}}UpdateBack struct {
	{{range .Fields}}{{.}}
    {{end}}
}

// Update ...
func (p *{{.StructName}}Update) Update() (*{{.StructName}}UpdateBack, error) {
	dao := &dao.{{.StructName}}Dao{}
	data, err := dao.Update{{.StructName}}(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back  {{.StructName}}UpdateBack 
	//这里是模板创建生成的，正常应该输入参数根据实际情况赋值给dao需要的参数
	tmpbb , _  :=  json.Marshal(*data)
	json.Unmarshal(tmpbb, &back)
	return &back, nil
}

`
