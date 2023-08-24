package mysql

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/dal"
	"github.com/keycasiter/3g_game/biz/model/po"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"strings"
)

func (WarBookDal) TableName() string {
	return "warbook"
}

type WarBookDal struct {
}

func NewWarbook() *WarBookDal {
	return &WarBookDal{}
}

func (g *WarBookDal) QueryWarbookList(ctx context.Context, condition *vo.QueryWarbookCondition) ([]*po.Warbook, error) {
	list := make([]*po.Warbook, 0)
	conn := dal.DataBase.Model(&po.Warbook{})

	//条件查询
	if condition.Id > 0 {
		conn.Where("id = ?", condition.Id)
	}
	if len(condition.Ids) > 0 {
		conn.Where("id in (?)", condition.Ids)
	}
	if strings.Trim(condition.Name, " ") != "" {
		conn.Where("name like ?", fmt.Sprintf("%%%s%%", condition.Name))
	}
	if condition.Level > 0 {
		conn.Where("level = ?", condition.Level)
	}

	if err := conn.Offset(condition.Offset).Limit(condition.Limit).Find(&list).Error; err != nil {
		hlog.CtxErrorf(ctx, "QueryWarbookList err:%v", err)
		return list, err
	}
	return list, nil
}
