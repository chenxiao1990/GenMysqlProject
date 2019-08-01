
package dao
import (
	"Out/model"
)

// HartDao ...
type HartDao struct {
}

// Create 增
func (*HartDao) Create(m *model.Hart) (*model.Hart, error) {
	err := model.DB.Create(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Delete 删
func (*HartDao) Delete(id int) error {
	err := model.DB.Delete(&model.Hart{ID: id}).Error
	return err
}

// SelectByID 查
func (*HartDao) SelectByID(id int) (*model.Hart, error) {

	var m model.Hart
	err := model.DB.Where("id = ?", id).Last(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// Update 改  map[string]interface{}{"name": "hello", "age": 18, "actived": false}
func (*HartDao) Update(id int, update map[string]interface{}) (*model.Hart, error) {

	var m model.Hart
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
