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

func (TacticDal) TableName() string {
	return "tactic"
}

type TacticDal struct {
}

func NewTactic() *TacticDal {
	return &TacticDal{}
}

func (g *TacticDal) QueryTacticList(ctx context.Context, condition *vo.QueryTacticCondition) ([]*po.Tactic, error) {
	list := make([]*po.Tactic, 0)
	conn := dal.DataBase.Model(&po.Tactic{})

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
	if condition.Quality > 0 {
		conn.Where("quality = ?", condition.Quality)
	}
	if condition.Source > 0 {
		conn.Where("source = ?", condition.Source)
	}
	if condition.Type > 0 {
		conn.Where("type = ?", condition.Type)
	}

	if err := conn.Find(&list).
		Offset(condition.Offset).
		Limit(condition.Limit).Error; err != nil {
		hlog.CtxErrorf(ctx, "QueryTacticList err:%v", err)
		return list, err
	}
	return list, nil
}
