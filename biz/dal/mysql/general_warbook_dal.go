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

func (GeneralWarBookDal) TableName() string {
	return "general_warbook"
}

type GeneralWarBookDal struct {
}

func NewGeneralWarbook() *GeneralWarBookDal {
	return &GeneralWarBookDal{}
}

func (g *GeneralWarBookDal) QueryGeneralWarbookList(ctx context.Context, condition *vo.QueryGeneralWarbookCondition) ([]*po.GeneralWarbook, error) {
	list := make([]*po.GeneralWarbook, 0)
	conn := dal.DataBase.Model(&po.GeneralWarbook{})

	//条件查询
	if condition.Id > 0 {
		conn.Where("id = ?", condition.Id)
	}
	if condition.GeneralID > 0 {
		conn.Where("general_id = ?", condition.GeneralID)
	}
	if strings.Trim(condition.Name, " ") != "" {
		conn.Where("name like ?", fmt.Sprintf("%%%s%%", condition.Name))
	}

	if err := conn.Offset(condition.OffSet).Limit(condition.Limit).Find(&list).Error; err != nil {
		hlog.CtxErrorf(ctx, "QueryGeneralWarbookList err:%v", err)
		return list, err
	}
	return list, nil
}
