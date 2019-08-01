package dao

var DapTemplate = `
package dao
import (
	"{{.ProjectName}}/model"
)

// {{.StructName}}Dao ...
type {{.StructName}}Dao struct {
}

// Create 增
func (*{{.StructName}}Dao) Create(m *model.{{.StructName}}) (*model.{{.StructName}}, error) {
	err := model.DB.Create(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Delete 删
func (*{{.StructName}}Dao) Delete(id int) error {
	err := model.DB.Delete(&model.{{.StructName}}{ID: id}).Error
	return err
}

// SelectByID 查
func (*{{.StructName}}Dao) SelectByID(id int) (*model.{{.StructName}}, error) {

	var m model.{{.StructName}}
	err := model.DB.Where("id = ?", id).Last(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// Update 改  map[string]interface{}{"name": "hello", "age": 18, "actived": false}
func (*{{.StructName}}Dao) Update(id int, update map[string]interface{}) (*model.{{.StructName}}, error) {

	var m model.{{.StructName}}
	err := model.DB.Model(&m).Where("id = ?", id).Updates(update).Error

	if err != nil {
		return nil, err
	}
	err = model.DB.Where("id = ?", id).Last(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}
`
