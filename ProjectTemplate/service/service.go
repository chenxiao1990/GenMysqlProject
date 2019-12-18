package service

var ServiceTemplate = `

package service
import (
	"{{.ProjectName}}/dao"
	"{{.ProjectName}}/model"
)

// {{.StructName}}Service 服务
type {{.StructName}}Service struct {
}

// {{.StructName}}CreateParam  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type {{.StructName}}CreateParam struct {
	model.{{.StructName}}
}

// {{.StructName}}CreateBack  返回参数
type {{.StructName}}CreateBack struct {
	model.{{.StructName}}
}

// Create 创建
func ( *{{.StructName}}Service) Create(p *{{.StructName}}CreateParam) (*{{.StructName}}CreateBack, error) {
	dao := &dao.{{.StructName}}Dao{}
	 
	data, err := dao.Create(&p.{{.StructName}})

	if err != nil {
		return nil, err

	}

	var back = {{.StructName}}CreateBack {
		*data,
	}
	  
	return &back, nil
}

 
// Delete  ...
func ( *{{.StructName}}Service) Delete(id int) error {
	dao := &dao.{{.StructName}}Dao{}
	return dao.Delete(id)
}

 

// {{.StructName}}SelectBack  返回参数
type {{.StructName}}SelectBack struct {
	model.{{.StructName}}
}

// Select ...
func (*{{.StructName}}Service) Select(id int) (*{{.StructName}}SelectBack, error) {
	dao := &dao.{{.StructName}}Dao{}
	data, err := dao.SelectByID(id)
	if err != nil {
		return nil, err
	}
	var back = {{.StructName}}SelectBack{
		*data,
	}
	  
	return &back, nil
}

// {{.StructName}}UpdateParam   参数 
type {{.StructName}}UpdateParam struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// {{.StructName}}UpdateBack  返回参数
type {{.StructName}}UpdateBack struct {
	model.{{.StructName}}
}

// Update ...
func (*{{.StructName}}Service) Update(p *{{.StructName}}UpdateParam) (*{{.StructName}}UpdateBack, error) {
	dao := &dao.{{.StructName}}Dao{}
	data, err := dao.Update(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back = {{.StructName}}UpdateBack {
		*data,
	}
	 
	return &back, nil
}

`
