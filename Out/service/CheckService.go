

package service
import (
	"Out/dao"
	"Out/model"
)

// CheckService 服务
type CheckService struct {
}

// CheckCreateParam  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type CheckCreateParam struct {
	model.Check
}

// CheckCreateBack  返回参数
type CheckCreateBack struct {
	model.Check
}

// Create 创建
func ( *CheckService) Create(p *CheckCreateParam) (*CheckCreateBack, error) {
	dao := &dao.CheckDao{}
	 
	data, err := dao.CreateCheck(&p.Check)

	if err != nil {
		return nil, err

	}

	var back = CheckCreateBack {
		*data,
	}
	  
	return &back, nil
}

// CheckDeleteParam   参数  
type CheckDeleteParam struct {
	ID int "json:\"id\" binding:\"required\""
}
// CheckDeleteBack  返回参数
type CheckDeleteBack struct {
	 
}
// Delete  ...
func ( *CheckService) Delete(p *CheckDeleteParam) error {
	dao := &dao.CheckDao{}
	return dao.DeleteCheck(p.ID)
}

// CheckSelectParam   参数 
type CheckSelectParam struct {
	ID int "json:\"id\" binding:\"required\""
}

// CheckSelectBack  返回参数
type CheckSelectBack struct {
	model.Check
}

// Select ...
func (*CheckService) Select(p *CheckSelectParam) (*CheckSelectBack, error) {
	dao := &dao.CheckDao{}
	data, err := dao.SelectCheckByID(p.ID)
	if err != nil {
		return nil, err
	}
	var back = CheckSelectBack{
		*data,
	}
	  
	return &back, nil
}

// CheckUpdateParam   参数 
type CheckUpdateParam struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// CheckUpdateBack  返回参数
type CheckUpdateBack struct {
	model.Check
}

// Update ...
func (*CheckService) Update(p *CheckUpdateParam) (*CheckUpdateBack, error) {
	dao := &dao.CheckDao{}
	data, err := dao.UpdateCheck(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back = CheckUpdateBack {
		*data,
	}
	 
	return &back, nil
}

