package mysql

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/model/po"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"strings"
)

func (SpecialTechDal) TableName() string {
	return "special_tech"
}

type SpecialTechDal struct {
}

func NewSpecialTech() *SpecialTechDal {
	return &SpecialTechDal{}
}

func (g *SpecialTechDal) QuerySpecialTechList(ctx context.Context, condition *vo.QuerySpecialTechCondition) ([]*po.SpecialTech, error) {
	list := make([]*po.SpecialTech, 0)
	conn := DataBase.Model(&po.SpecialTech{})

	//条件查询
	if condition.Id > 0 {
		conn.Where("id = ?", condition.Id)
	}
	if condition.Type > 0 {
		conn.Where("type = ?", condition.Type)
	}
	if condition.Level > 0 {
		conn.Where("level = ?", condition.Level)
	}
	if strings.Trim(condition.Name, " ") != "" {
		conn.Where("name like ?", fmt.Sprintf("%%%s%%", condition.Name))
	}

	if err := conn.Offset(condition.Offset).Limit(condition.Limit).Find(&list).Error; err != nil {
		hlog.CtxErrorf(ctx, "QuerySpecialTechList err:%v", err)
		return list, err
	}
	return list, nil
}
