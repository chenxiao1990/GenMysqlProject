

package service
import (
	"Out/dao"
	"Out/model"
	"encoding/json"
)

// HealthCreate  服务：创建  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type HealthCreate struct {
	ID int `json:"id"`
    UID int `json:"uid"`
    HealStep int `json:"heal_step"`
    Power int `json:"power"`
    Cal float64 `json:"cal"`
    Created int `json:"created"`
    
}

// HealthCreateBack  返回参数
type HealthCreateBack struct {
	ID int `json:"id"`
    UID int `json:"uid"`
    HealStep int `json:"heal_step"`
    Power int `json:"power"`
    Cal float64 `json:"cal"`
    Created int `json:"created"`
    
}

// Create 创建
func (p *HealthCreate) Create() (*HealthCreateBack, error) {
	dao := &dao.HealthDao{}
	var model  model.Health 
	//这里是模板创建生成的，正常应该输入参数根据实际情况赋值给dao需要的参数
	tmpbb , _ :=  json.Marshal(*p)
	json.Unmarshal(tmpbb, &model)
	data, err := dao.CreateHealth(&model)

	if err != nil {
		return nil, err

	}

	var back HealthCreateBack
	//这里是模板创建生成的，正常应该输入参数根据实际情况赋值给dao需要的参数
	tmpbb , _  =  json.Marshal(*data)
	json.Unmarshal(tmpbb, &back)
 
	return &back, nil
}

// HealthDelete  服务：删除   
type HealthDelete struct {
	ID int "json:\"id\" binding:\"required\""
}

// Delete 创建
func (p *HealthDelete) Delete() error {
	dao := &dao.HealthDao{}
	return dao.DeleteHealth(p.ID)
}

// HealthSelect  服务：查询
type HealthSelect struct {
	ID int "json:\"id\" binding:\"required\""
}

// HealthSelectBack  返回参数
type HealthSelectBack struct {
	ID int `json:"id"`
    UID int `json:"uid"`
    HealStep int `json:"heal_step"`
    Power int `json:"power"`
    Cal float64 `json:"cal"`
    Created int `json:"created"`
    
}

// Select ...
func (p *HealthSelect) Select() (*HealthSelectBack, error) {
	dao := &dao.HealthDao{}
	data, err := dao.SelectHealthByID(p.ID)
	if err != nil {
		return nil, err
	}
	var back HealthSelectBack 
	//这里是模板创建生成的，正常应该输入参数根据实际情况赋值给dao需要的参数
	tmpbb , _  :=  json.Marshal(*data)
	json.Unmarshal(tmpbb, &back)

	return &back, nil
}

// HealthUpdate  服务：更新
type HealthUpdate struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// HealthUpdateBack  返回参数
type HealthUpdateBack struct {
	ID int `json:"id"`
    UID int `json:"uid"`
    HealStep int `json:"heal_step"`
    Power int `json:"power"`
    Cal float64 `json:"cal"`
    Created int `json:"created"`
    
}

// Update ...
func (p *HealthUpdate) Update() (*HealthUpdateBack, error) {
	dao := &dao.HealthDao{}
	data, err := dao.UpdateHealth(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back  HealthUpdateBack 
	//这里是模板创建生成的，正常应该输入参数根据实际情况赋值给dao需要的参数
	tmpbb , _  :=  json.Marshal(*data)
	json.Unmarshal(tmpbb, &back)
	return &back, nil
}

