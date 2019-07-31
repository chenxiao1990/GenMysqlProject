

package service
import (
	"Out/dao"
	"Out/model"
)

// HartService 服务
type HartService struct {
}

// HartCreateParam  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type HartCreateParam struct {
	model.Hart
}

// HartCreateBack  返回参数
type HartCreateBack struct {
	model.Hart
}

// Create 创建
func ( *HartService) Create(p *HartCreateParam) (*HartCreateBack, error) {
	dao := &dao.HartDao{}
	 
	data, err := dao.CreateHart(&p.Hart)

	if err != nil {
		return nil, err

	}

	var back = HartCreateBack {
		*data,
	}
	  
	return &back, nil
}

// HartDeleteParam   参数  
type HartDeleteParam struct {
	ID int "json:\"id\" binding:\"required\""
}
// HartDeleteBack  返回参数
type HartDeleteBack struct {
	 
}
// Delete  ...
func ( *HartService) Delete(p *HartDeleteParam) error {
	dao := &dao.HartDao{}
	return dao.DeleteHart(p.ID)
}

// HartSelectParam   参数 
type HartSelectParam struct {
	ID int "json:\"id\" binding:\"required\""
}

// HartSelectBack  返回参数
type HartSelectBack struct {
	model.Hart
}

// Select ...
func (*HartService) Select(p *HartSelectParam) (*HartSelectBack, error) {
	dao := &dao.HartDao{}
	data, err := dao.SelectHartByID(p.ID)
	if err != nil {
		return nil, err
	}
	var back = HartSelectBack{
		*data,
	}
	  
	return &back, nil
}

// HartUpdateParam   参数 
type HartUpdateParam struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// HartUpdateBack  返回参数
type HartUpdateBack struct {
	model.Hart
}

// Update ...
func (*HartService) Update(p *HartUpdateParam) (*HartUpdateBack, error) {
	dao := &dao.HartDao{}
	data, err := dao.UpdateHart(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back = HartUpdateBack {
		*data,
	}
	 
	return &back, nil
}

