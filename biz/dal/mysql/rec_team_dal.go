package mysql

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/dal"
	"github.com/keycasiter/3g_game/biz/model/po"
	"github.com/keycasiter/3g_game/biz/model/vo"
)

func (RecTeamDal) TableName() string {
	return "rec_team"
}

type RecTeamDal struct {
}

func NewRecTeam() *RecTeamDal {
	return &RecTeamDal{}
}

func (g *RecTeamDal) QueryRecTeamList(ctx context.Context, condition *vo.QueryRecTeamCondition) ([]*po.RecTeam, error) {
	list := make([]*po.RecTeam, 0)
	conn := dal.DataBase.Model(&po.RecTeam{})

	//条件查询
	if condition.Name != "" {
		conn.Where("name like ?", fmt.Sprintf("%%%s%%", condition.Name))
	}
	if condition.Group > 0 {
		conn.Where("`group` = ?", condition.Group)
	}

	if err := conn.Offset(condition.Offset).Limit(condition.Limit).Find(&list).Error; err != nil {
		hlog.CtxErrorf(ctx, "QueryRecTeamList err:%v", err)
		return list, err
	}
	return list, nil
}
