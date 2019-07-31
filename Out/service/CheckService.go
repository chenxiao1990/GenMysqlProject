

package service
import (
	"Out/dao"
	"Out/model"
	"encoding/json"
)

// CheckCreate  服务：创建  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type CheckCreate struct {
	ID int `json:"id"`
    UID int `json:"uid"`
    AreaID string `json:"area_id"`
    GradeID int `json:"grade_id"`
    ClassID int `json:"class_id"`
    Created int `json:"created"`
    Updated int `json:"updated"`
    
}

// CheckCreateBack  返回参数
type CheckCreateBack struct {
	ID int `json:"id"`
    UID int `json:"uid"`
    AreaID string `json:"area_id"`
    GradeID int `json:"grade_id"`
    ClassID int `json:"class_id"`
    Created int `json:"created"`
    Updated int `json:"updated"`
    
}

// Create 创建
func (p *CheckCreate) Create() (*CheckCreateBack, error) {
	dao := &dao.CheckDao{}
	var model  model.Check 
	//这里是模板创建生成的，正常应该输入参数根据实际情况赋值给dao需要的参数
	tmpbb , _ :=  json.Marshal(*p)
	json.Unmarshal(tmpbb, &model)
	data, err := dao.CreateCheck(&model)

	if err != nil {
		return nil, err

	}

	var back CheckCreateBack
	//这里是模板创建生成的，正常应该输入参数根据实际情况赋值给dao需要的参数
	tmpbb , _  =  json.Marshal(*data)
	json.Unmarshal(tmpbb, &back)
 
	return &back, nil
}

// CheckDelete  服务：删除   
type CheckDelete struct {
	ID int "json:\"id\" binding:\"required\""
}

// Delete 创建
func (p *CheckDelete) Delete() error {
	dao := &dao.CheckDao{}
	return dao.DeleteCheck(p.ID)
}

// CheckSelect  服务：查询
type CheckSelect struct {
	ID int "json:\"id\" binding:\"required\""
}

// CheckSelectBack  返回参数
type CheckSelectBack struct {
	ID int `json:"id"`
    UID int `json:"uid"`
    AreaID string `json:"area_id"`
    GradeID int `json:"grade_id"`
    ClassID int `json:"class_id"`
    Created int `json:"created"`
    Updated int `json:"updated"`
    
}

// Select ...
func (p *CheckSelect) Select() (*CheckSelectBack, error) {
	dao := &dao.CheckDao{}
	data, err := dao.SelectCheckByID(p.ID)
	if err != nil {
		return nil, err
	}
	var back CheckSelectBack 
	//这里是模板创建生成的，正常应该输入参数根据实际情况赋值给dao需要的参数
	tmpbb , _  :=  json.Marshal(*data)
	json.Unmarshal(tmpbb, &back)

	return &back, nil
}

// CheckUpdate  服务：更新
type CheckUpdate struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// CheckUpdateBack  返回参数
type CheckUpdateBack struct {
	ID int `json:"id"`
    UID int `json:"uid"`
    AreaID string `json:"area_id"`
    GradeID int `json:"grade_id"`
    ClassID int `json:"class_id"`
    Created int `json:"created"`
    Updated int `json:"updated"`
    
}

// Update ...
func (p *CheckUpdate) Update() (*CheckUpdateBack, error) {
	dao := &dao.CheckDao{}
	data, err := dao.UpdateCheck(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back  CheckUpdateBack 
	//这里是模板创建生成的，正常应该输入参数根据实际情况赋值给dao需要的参数
	tmpbb , _  :=  json.Marshal(*data)
	json.Unmarshal(tmpbb, &back)
	return &back, nil
}

