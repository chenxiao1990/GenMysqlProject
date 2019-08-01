

package service
import (
	"Out/dao"
	"Out/model"
)

// HealthService 服务
type HealthService struct {
}

// HealthCreateParam  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type HealthCreateParam struct {
	model.Health
}

// HealthCreateBack  返回参数
type HealthCreateBack struct {
	model.Health
}

// Create 创建
func ( *HealthService) Create(p *HealthCreateParam) (*HealthCreateBack, error) {
	dao := &dao.HealthDao{}
	 
	data, err := dao.Create(&p.Health)

	if err != nil {
		return nil, err

	}

	var back = HealthCreateBack {
		*data,
	}
	  
	return &back, nil
}

// HealthDeleteParam   参数  
type HealthDeleteParam struct {
	ID int "json:\"id\" binding:\"required\""
}
// HealthDeleteBack  返回参数
type HealthDeleteBack struct {
	 
}
// Delete  ...
func ( *HealthService) Delete(p *HealthDeleteParam) error {
	dao := &dao.HealthDao{}
	return dao.Delete(p.ID)
}

// HealthSelectParam   参数 
type HealthSelectParam struct {
	ID int "json:\"id\" binding:\"required\""
}

// HealthSelectBack  返回参数
type HealthSelectBack struct {
	model.Health
}

// Select ...
func (*HealthService) Select(p *HealthSelectParam) (*HealthSelectBack, error) {
	dao := &dao.HealthDao{}
	data, err := dao.SelectByID(p.ID)
	if err != nil {
		return nil, err
	}
	var back = HealthSelectBack{
		*data,
	}
	  
	return &back, nil
}

// HealthUpdateParam   参数 
type HealthUpdateParam struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// HealthUpdateBack  返回参数
type HealthUpdateBack struct {
	model.Health
}

// Update ...
func (*HealthService) Update(p *HealthUpdateParam) (*HealthUpdateBack, error) {
	dao := &dao.HealthDao{}
	data, err := dao.Update(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back = HealthUpdateBack {
		*data,
	}
	 
	return &back, nil
}

