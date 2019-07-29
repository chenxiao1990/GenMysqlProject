package dao

import (
	"Out/model"
)

// UserDao ...
type UserDao struct {
}

// CreateUser 增
func (*UserDao) CreateUser(m *model.User) (*model.User, error) {
	err := model.DB.Create(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// DeleteUser 删
func (*UserDao) DeleteUser(id int) error {
	err := model.DB.Delete(&model.User{ID: id}).Error
	return err
}

// SelectUserByID 查
func (*UserDao) SelectUserByID(id int) (*model.User, error) {

	var m model.User
	err := model.DB.Where("id = ?", id).Last(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// UpdateUser 改  map[string]interface{}{"name": "hello", "age": 18, "actived": false}
func (*UserDao) UpdateUser(id int, update map[string]interface{}) (*model.User, error) {

	var m model.User
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
