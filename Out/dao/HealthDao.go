
package dao
import (
	"Out/model"
)

// HealthDao ...
type HealthDao struct {
}

// Create 增
func (*HealthDao) Create(m *model.Health) (*model.Health, error) {
	err := model.DB.Create(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Delete 删
func (*HealthDao) Delete(id int) error {
	err := model.DB.Delete(&model.Health{ID: id}).Error
	return err
}

// SelectByID 查
func (*HealthDao) SelectByID(id int) (*model.Health, error) {

	var m model.Health
	err := model.DB.Where("id = ?", id).Last(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// Update 改  map[string]interface{}{"name": "hello", "age": 18, "actived": false}
func (*HealthDao) Update(id int, update map[string]interface{}) (*model.Health, error) {

	var m model.Health
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
