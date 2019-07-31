

package service
import (
	"Out/dao"
	"Out/model"
	"encoding/json"
)

// HartCreate  服务：创建  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type HartCreate struct {
	ID int `json:"id"`
    HealHr int `json:"heal_hr"`
    Created int `json:"created"`
    Updated int `json:"updated"`
    
}

// HartCreateBack  返回参数
type HartCreateBack struct {
	ID int `json:"id"`
    HealHr int `json:"heal_hr"`
    Created int `json:"created"`
    Updated int `json:"updated"`
    
}

// Create 创建
func (p *HartCreate) Create() (*HartCreateBack, error) {
	dao := &dao.HartDao{}
	var model  model.Hart 
	//这里是模板创建生成的，正常应该输入参数根据实际情况赋值给dao需要的参数
	tmpbb , _ :=  json.Marshal(*p)
	json.Unmarshal(tmpbb, &model)
	data, err := dao.CreateHart(&model)

	if err != nil {
		return nil, err

	}

	var back HartCreateBack
	//这里是模板创建生成的，正常应该输入参数根据实际情况赋值给dao需要的参数
	tmpbb , _  =  json.Marshal(*data)
	json.Unmarshal(tmpbb, &back)
 
	return &back, nil
}

// HartDelete  服务：删除   
type HartDelete struct {
	ID int "json:\"id\" binding:\"required\""
}

// Delete 创建
func (p *HartDelete) Delete() error {
	dao := &dao.HartDao{}
	return dao.DeleteHart(p.ID)
}

// HartSelect  服务：查询
type HartSelect struct {
	ID int "json:\"id\" binding:\"required\""
}

// HartSelectBack  返回参数
type HartSelectBack struct {
	ID int `json:"id"`
    HealHr int `json:"heal_hr"`
    Created int `json:"created"`
    Updated int `json:"updated"`
    
}

// Select ...
func (p *HartSelect) Select() (*HartSelectBack, error) {
	dao := &dao.HartDao{}
	data, err := dao.SelectHartByID(p.ID)
	if err != nil {
		return nil, err
	}
	var back HartSelectBack 
	//这里是模板创建生成的，正常应该输入参数根据实际情况赋值给dao需要的参数
	tmpbb , _  :=  json.Marshal(*data)
	json.Unmarshal(tmpbb, &back)

	return &back, nil
}

// HartUpdate  服务：更新
type HartUpdate struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// HartUpdateBack  返回参数
type HartUpdateBack struct {
	ID int `json:"id"`
    HealHr int `json:"heal_hr"`
    Created int `json:"created"`
    Updated int `json:"updated"`
    
}

// Update ...
func (p *HartUpdate) Update() (*HartUpdateBack, error) {
	dao := &dao.HartDao{}
	data, err := dao.UpdateHart(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back  HartUpdateBack 
	//这里是模板创建生成的，正常应该输入参数根据实际情况赋值给dao需要的参数
	tmpbb , _  :=  json.Marshal(*data)
	json.Unmarshal(tmpbb, &back)
	return &back, nil
}

