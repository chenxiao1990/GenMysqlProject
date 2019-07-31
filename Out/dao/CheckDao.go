
package dao
import (
	"Out/model"
)

// CheckDao ...
type CheckDao struct {
}

// CreateCheck 增
func (*CheckDao) CreateCheck(m *model.Check) (*model.Check, error) {
	err := model.DB.Create(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// DeleteCheck 删
func (*CheckDao) DeleteCheck(id int) error {
	err := model.DB.Delete(&model.Check{ID: id}).Error
	return err
}

// SelectCheckByID 查
func (*CheckDao) SelectCheckByID(id int) (*model.Check, error) {

	var m model.Check
	err := model.DB.Where("id = ?", id).Last(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// UpdateCheck 改  map[string]interface{}{"name": "hello", "age": 18, "actived": false}
func (*CheckDao) UpdateCheck(id int, update map[string]interface{}) (*model.Check, error) {

	var m model.Check
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
