package service

import (
	"Out/dao"
	"Out/model"
	"encoding/json"
)

// UserCreate  服务：创建  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type UserCreate struct {
	ID int `json:"id"`
    UID int `json:"uid"`
    Name string `json:"name"`
    Phone string `json:"phone"`
    Pass string `json:"pass"`
    Status int `json:"status"`
    
}

// UserCreateBack  返回参数
type UserCreateBack struct {
	ID int `json:"id"`
    UID int `json:"uid"`
    Name string `json:"name"`
    Phone string `json:"phone"`
    Pass string `json:"pass"`
    Status int `json:"status"`
    
}

// Create 创建
func (p *UserCreate) Create() (*UserCreateBack, error) {
	dao := &dao.UserDao{}
	var model  model.User 
	//这里是模板创建生成的，正常应该输入参数根据实际情况赋值给dao需要的参数
	tmpbb , _ :=  json.Marshal(*p)
	json.Unmarshal(tmpbb, &model)
	data, err := dao.CreateUser(&model)

	if err != nil {
		return nil, err

	}

	var back UserCreateBack
	//这里是模板创建生成的，正常应该输入参数根据实际情况赋值给dao需要的参数
	tmpbb , _  =  json.Marshal(*data)
	json.Unmarshal(tmpbb, &back)
 
	return &back, nil
}

// UserDelete  服务：删除   
type UserDelete struct {
	ID int `json:"id" binding:"required"`
}

// Delete 创建
func (p *UserDelete) Delete() error {
	dao := &dao.UserDao{}
	return dao.DeleteUser(p.ID)
}

// UserSelect  服务：查询
type UserSelect struct {
	ID int `json:"id" binding:"required"`
}

// UserSelectBack  返回参数
type UserSelectBack struct {
	ID int `json:"id"`
    UID int `json:"uid"`
    Name string `json:"name"`
    Phone string `json:"phone"`
    Pass string `json:"pass"`
    Status int `json:"status"`
    
}

// Select ...
func (p *UserSelect) Select() (*UserSelectBack, error) {
	dao := &dao.UserDao{}
	data, err := dao.SelectUserByID(p.ID)
	if err != nil {
		return nil, err
	}
	var back UserSelectBack 
	//这里是模板创建生成的，正常应该输入参数根据实际情况赋值给dao需要的参数
	tmpbb , _  :=  json.Marshal(*data)
	json.Unmarshal(tmpbb, &back)

	return &back, nil
}

// UserUpdate  服务：更新
type UserUpdate struct {
	ID    int                    `json:"id" binding:"required"`
	Param map[string]interface{} `json:"param" binding:"required"`
}

// UserUpdateBack  返回参数
type UserUpdateBack struct {
	ID int `json:"id"`
    UID int `json:"uid"`
    Name string `json:"name"`
    Phone string `json:"phone"`
    Pass string `json:"pass"`
    Status int `json:"status"`
    
}

// Update ...
func (p *UserUpdate) Update() (*UserUpdateBack, error) {
	dao := &dao.UserDao{}
	data, err := dao.UpdateUser(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back  UserUpdateBack 
	//这里是模板创建生成的，正常应该输入参数根据实际情况赋值给dao需要的参数
	tmpbb , _  :=  json.Marshal(*data)
	json.Unmarshal(tmpbb, &back)
	return &back, nil
}
