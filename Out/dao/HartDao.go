
package dao
import (
	"Out/model"
)

// HartDao ...
type HartDao struct {
}

// CreateHart 增
func (*HartDao) CreateHart(m *model.Hart) (*model.Hart, error) {
	err := model.DB.Create(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// DeleteHart 删
func (*HartDao) DeleteHart(id int) error {
	err := model.DB.Delete(&model.Hart{ID: id}).Error
	return err
}

// SelectHartByID 查
func (*HartDao) SelectHartByID(id int) (*model.Hart, error) {

	var m model.Hart
	err := model.DB.Where("id = ?", id).Last(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// UpdateHart 改  map[string]interface{}{"name": "hello", "age": 18, "actived": false}
func (*HartDao) UpdateHart(id int, update map[string]interface{}) (*model.Hart, error) {

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
